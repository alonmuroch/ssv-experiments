package ssv_chain

import (
	"fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	"ssv-experiments/ssv_chain/types"
)

type App struct {
	State *types.State
}

// ProcessBlock is a CometBFT compatible request that is called during FinalizeBlock.
//
//	It applies the block to the state and returns a CometBFT compatible response
func (app *App) ProcessBlock(req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	if err := app.ValidateProposer(req.ProposerAddress); err != nil {
		return nil, err
	}

	return app.ApplyBlock(req)
}

func (app *App) ApplyBlock(req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	panic("implement")
}

func (app *App) ValidateProposer(address []byte) error {
	panic("implement")
}

func (app *App) ApplyTransactions(txs [][]byte) error {
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
