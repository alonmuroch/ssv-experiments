package ssv_chain

import (
	"bytes"
	"fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	v12 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	"ssv-experiments/ssv_chain/common"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/types"
)

const (
	TransactionErrorCodeSpace   = "tx_error"
	TransactionSuccessCodeSpace = ""
	TransactionSuccessCode      = 0
	TransactionErrorCode        = 1 // according to cosmos SDK a successful tx code is 0
)

type App struct {
	State  *types.State
	config *types.Configure
}

// FinalizeBlock is a CometBFT compatible ABCI request.
//
//	It applies the block to the state and returns a CometBFT compatible response
func (app *App) FinalizeBlock(req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	txResults, err := app.ApplyBlock(req)
	if err != nil {
		return nil, err
	}

	valUpdates, err := app.ApplyValidatorRewardsPenalties(req.Misbehavior, req.DecidedLastCommit)

	// calculate new state hash
	h, err := app.State.HashTreeRoot()
	if err != nil {
		return nil, err
	}

	return &v1.FinalizeBlockResponse{
		TxResults:        txResults,
		ValidatorUpdates: valUpdates,
		AppHash:          h[:],
	}, nil
}

func (app *App) ApplyValidatorRewardsPenalties(misbehavior []v1.Misbehavior, commits v1.CommitInfo) ([]v1.ValidatorUpdate, error) {
	ret := make([]v1.ValidatorUpdate, 0)

	// slash validators for misbehaviours
	for _, m := range misbehavior {
		switch m.Type {
		case v1.MISBEHAVIOR_TYPE_DUPLICATE_VOTE:
			val := app.State.ValidatorByAddress(m.Validator.Address)
			if val == nil {
				return nil, fmt.Errorf("valdiator not found")
			}

			ret = append(ret, v1.ValidatorUpdate{
				Power:       0,
				PubKeyBytes: val.PublicKey.Key,
				PubKeyType:  "", // TODO
			})

			if err := val.Slash(app.State, app.config); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown misbehavirou")
		}
	}

	// penalize validators for misbehaviour
	for _, v := range commits.Votes {
		if v.BlockIdFlag != v12.BlockIDFlagCommit {
			val := app.State.ValidatorByAddress(v.Validator.Address)
			if val == nil {
				return nil, fmt.Errorf("valdiator not found")
			}

			ret = append(ret, v1.ValidatorUpdate{
				Power:       v.Validator.Power - 1,
				PubKeyBytes: val.PublicKey.Key,
				PubKeyType:  "", // TODO
			})

			if err := val.Penalize(app.State, app.config, app.config.MissedValidationPenalty); err != nil {
				return nil, err
			}
		}
	}

	return ret, nil
}

func (app *App) ApplyBlock(req *v1.FinalizeBlockRequest) ([]*v1.ExecTxResult, error) {
	if err := app.ApplyBlockHeight(app.State, uint64(req.Height)); err != nil {
		return nil, err
	}
	if err := app.ValidateProposer(uint64(req.Height), req.ProposerAddress); err != nil {
		return nil, err
	}

	app.State.LatestBlockHeaderHash = req.Hash

	postState, results, err := app.ApplyTransactions(req.Txs)
	if err != nil {
		return nil, err
	}
	app.State = postState

	if err := app.UpdateBalancesPerFee(); err != nil {
		return nil, err
	}

	return results, nil
}

func (app *App) ValidateProposer(height uint64, address []byte) error {
	panic("implement")
}

func (app *App) ApplyTransactions(txsRaw [][]byte) (*types.State, []*v1.ExecTxResult, error) {
	txs := make([]*types.Transaction, len(txsRaw))
	for _, txRaw := range txsRaw {
		tx := &types.SignedTransaction{}
		if err := tx.UnmarshalSSZ(txRaw); err != nil {
			return nil, nil, err
		}

		if bytes.Equal(tx.Transaction.Address, app.config.SystemTxSigner) { // system tx
			if err := app.ValidateSystemTransaction(tx); err != nil {
				return nil, nil, err
			}
		} else { // user transaction
			acc := app.State.AccountByAddress(tx.Transaction.Address)
			if acc == nil {
				return nil, nil, fmt.Errorf("account not found")
			}

			if !common.VerifySignature(acc.Network, &tx.Transaction, tx.Signature, tx.Transaction.Address) {
				return nil, nil, fmt.Errorf("invalid signature")
			}
		}
	}

	ctx := &operations.Context{
		Config: app.config,
		State:  app.State.DeepCopy(),
	}

	rec := ProcessTransactions(ctx, app.State, txs)
	ret := make([]*v1.ExecTxResult, len(txsRaw))
	for i := range txsRaw {
		exeRes := &v1.ExecTxResult{
			GasUsed:   int64(rec.GasPerTx[i]),
			Codespace: TransactionSuccessCodeSpace,
			Code:      TransactionSuccessCode,
		}
		if rec.Errors[i] != nil {
			exeRes.Log = rec.Errors[i].Error()
			exeRes.Codespace = TransactionErrorCodeSpace
			exeRes.Code = TransactionErrorCode
		}

		ret[i] = exeRes
	}

	return ctx.State, ret, nil
}

// UpdateBalancesPerFee will apply fee calculation across all accounts
func (app *App) UpdateBalancesPerFee() error {
	for _, c := range app.State.Clusters {
		billableAccount := app.State.AccountByAddress(c.Address)
		if !c.Active {
			continue
		}
	clusterLoop:
		for _, inst := range c.Instances {
			operatorAccounts := app.State.OperatorAccountsByID(c.Operators)

			for i, priceTierIndex := range inst.PriceTierIndexes {
				op := app.State.Operators[i]
				priceTier := op.Tiers[priceTierIndex]
				if err := billableAccount.ReduceBalance(priceTier.Price, priceTier.PayableTokenAddress, priceTier.Network); err != nil {
					// if billable account can't be charge, mark cluster as not active and break.
					// Operators will not get paid for this block, but will stop executing for this cluster
					c.Active = false
					break clusterLoop // break instances loop
				}
				operatorAccounts[i].AddBalance(priceTier.Price, priceTier.PayableTokenAddress, priceTier.Network)
			}
		}
	}
	return nil
}

func (app *App) ApplyBlockHeight(state *types.State, newHeight uint64) error {
	if state.BlockHeight+1 != newHeight {
		return fmt.Errorf("invalid height")
	}
	state.BlockHeight++
	return nil
}

func (app *App) PenalizeMisbehaviour(misbehavior []v1.Misbehavior) error {
	panic("implement")
}

func (app *App) ValidateSystemTransaction(tx *types.SignedTransaction) error {
	panic("implement")
}
