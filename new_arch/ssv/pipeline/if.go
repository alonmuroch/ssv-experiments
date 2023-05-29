package pipeline

import (
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

// ContinueIfConsensusMessage execute the next item if first object is a QBFT Signed Message, otherwise skip next
func ContinueIfConsensusMessage(next PipelineF) PipelineF {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		if _, isConsensusMessage := objects[0].(*qbft.SignedMessage); isConsensusMessage {
			return next(runner, objects)
		}
		return nil, []interface{}{SkipNext}
	}
}

// ContinueIfPostConsensusMessage execute the next item if first object is a post consensus Message, otherwise skip next
func ContinueIfPostConsensusMessage(next PipelineF) PipelineF {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		if m, isPartialSigMessage := objects[0].(*types.SignedPartialSignatureMessages); isPartialSigMessage && m.Message.Type == types.PostConsensusPartialSig {
			return next(runner, objects)
		}
		return nil, []interface{}{SkipNext}
	}
}

// ContinueIfPreConsensusMessage execute the next item if first object is a pre consensus Message, otherwise skip next
func ContinueIfPreConsensusMessage(next PipelineF) PipelineF {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		isPreMsgType := func(t types.PartialSigMsgType) bool {
			if t == types.RandaoPartialSig || t == types.ValidatorRegistrationPartialSig || t == types.ContributionProofs || t == types.SelectionProofPartialSig {
				return true
			}
			return false
		}

		if m, isPartialSigMessage := objects[0].(*types.SignedPartialSignatureMessages); isPartialSigMessage && isPreMsgType(m.Message.Type) {
			return next(runner, objects)
		}
		return nil, []interface{}{SkipNext}
	}
}
