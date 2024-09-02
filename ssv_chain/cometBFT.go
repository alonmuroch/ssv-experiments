package ssv_chain

import (
	"fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	v12 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	"ssv-experiments/ssv_chain/types"
)

type App struct {
	State  *types.State
	config types.Configure
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
				PubKeyBytes: val.PublicKey.PK,
				PubKeyType:  "", // TODO
			})

			if err := val.Slash(app.State, app.config); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown misbehavirou")
		}
	}

	// penalize validators for misbehaviours
	for _, v := range commits.Votes {
		if v.BlockIdFlag != v12.BlockIDFlagCommit {
			val := app.State.ValidatorByAddress(v.Validator.Address)
			if val == nil {
				return nil, fmt.Errorf("valdiator not found")
			}

			ret = append(ret, v1.ValidatorUpdate{
				Power:       v.Validator.Power - 1,
				PubKeyBytes: val.PublicKey.PK,
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

	results, err := app.ApplyTransactions(req.Txs)
	if err != nil {
		return nil, err
	}

	if err := app.UpdateBalancesPerFee(); err != nil {
		return nil, err
	}

	return results, nil
}

func (app *App) ValidateProposer(height uint64, address []byte) error {
	panic("implement")
}

func (app *App) ApplyTransactions(txs [][]byte) ([]*v1.ExecTxResult, error) {
	panic("implement")
}

// UpdateBalancesPerFee will apply fee calculation across all accounts
func (app *App) UpdateBalancesPerFee() error {
	panic("implement")
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
