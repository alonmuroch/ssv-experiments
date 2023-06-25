package pipeline

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

type ControlSymbols uint64

const (
	InitPhase  = "InitPhase"
	StartPhase = "StartPhase"
	EndPhase   = "EndPhase"
)

const (
	// Stop means stop pipeline
	Stop ControlSymbols = iota
	// SkipNext means skip the next pipeline element
	SkipNext
	// SkipToPhase skips to a specific phase. objects[0] control symbol, objects[1] next phase, objects[2:] objects to pass to next phase
	SkipToPhase
)

// PipelineF is a function taking a runner and multiple objects to process a single action
// Those functions suppose to be chained together to produce a whole working message process for a runner
type PipelineF func(pipeline *Pipeline, objects ...interface{}) (error, []interface{})

type Pipeline struct {
	Runner   *ssv.Runner
	Instance *qbft.Instance
	Items    []PipelineF
	Phase    map[string]int // maps phase name to index
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Items: []PipelineF{},
		Phase: map[string]int{},
	}
}

func (p *Pipeline) Init() error {
	err, _ := Run(p, InitPhase, nil)
	return err
}

// ProcessMessage inputs a P2P message and passes it through the pipeline
func (p *Pipeline) ProcessMessage(msg interface{}) (error, []interface{}) {
	return Run(p, StartPhase, []interface{}{msg})
}

// Add a pipeline item
func (p *Pipeline) Add(f PipelineF) *Pipeline {
	p.Items = append(p.Items, f)
	return p
}

// MarkPhase marks a phase by name in the pipeline that can be jumped to
func (p *Pipeline) MarkPhase(name string) *Pipeline {
	p.Phase[name] = len(p.Items)
	return p
}

func (p *Pipeline) Stop() *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		return nil, []interface{}{Stop}
	})
	return p
}

func (p *Pipeline) StopINoPreConsensusQuorum() *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		// TODO check pre-consensus quorum
		if false {
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

func (p *Pipeline) StopIfNotDecided() *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		if !pipeline.Instance.Decided() {
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

// StopIfNoPartialSigQuorum checks if msg container for type has quorum, if not stop
func (p *Pipeline) StopIfNoPartialSigQuorum(t types.PartialSigMsgType) *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		// get container by type from runner
		// check quorum
		if false { // no quorum
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

// ValidateConsensusMessage validates consensus message (type, struct, etc), returns error if not valid
func (p *Pipeline) ValidateConsensusMessage() *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		_, ok := objects[0].(*qbft.SignedMessage)
		if ok {
			return nil, objects
		}
		return errors.New("not a consensus message"), nil
	})
	return p
}

// SkipIfNotQBFTMessageType will validate message type, will skip if not
func (p *Pipeline) SkipIfNotQBFTMessageType(nextPhase string, msgType uint64) *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		msg, ok := objects[0].(*qbft.SignedMessage)
		if !ok {
			return nil, append(
				[]interface{}{
					SkipToPhase,
					nextPhase,
				}, objects...)
		}

		if msg.Message.MsgType == msgType {
			return nil, objects
		}
		return nil, append(
			[]interface{}{
				SkipToPhase,
				nextPhase,
			}, objects...)
	})
	return p
}

func (p *Pipeline) SkipIfNotPreConsensusMessage(nextPhase string) *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		// check if pre consensus message

		if true { // consensus message
			return nil, objects
		}
		return nil, append(
			[]interface{}{
				SkipToPhase,
				nextPhase,
			}, objects...)
	})
	return p
}

// StopIfNotPostConsensusMessage will stop the pipeline if message is not post consensus message
func (p *Pipeline) StopIfNotPostConsensusMessage() *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		// check if post consensus message

		if true { // consensus message
			return nil, objects
		}
		return nil, append(
			[]interface{}{
				Stop,
			}, objects...)
	})
	return p
}

// IndexForPhase returns the index for a marked phase or error
func (p *Pipeline) IndexForPhase(phase string) (int, error) {
	if val, found := p.Phase[phase]; found {
		return val, nil
	} else {
		return 0, errors.New("phase not found")
	}
}
