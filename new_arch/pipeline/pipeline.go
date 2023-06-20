package pipeline

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
	"ssv-experiments/new_arch/types"
)

type ControlSymbols uint64

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

// ProcessMessage inputs a P2P message and passes it through the pipeline
func (p *Pipeline) ProcessMessage(msg interface{}) (error, []interface{}) {
	initObjs := []interface{}{msg}
	for i := range p.Items {
		f := p.Items[i]
		// check control symbols
		if len(initObjs) > 0 {
			if c, isControlSymbol := initObjs[0].(ControlSymbols); isControlSymbol {
				switch c {
				case Stop:
					return nil, []interface{}{}
				case SkipNext:
					continue
				case SkipToPhase:
					phaseName := initObjs[1].(string)
					if val, found := p.Phase[phaseName]; found {
						i = val // bump i to phase index
						initObjs = initObjs[2:]
						continue
					} else {
						return errors.New("phase not found"), []interface{}{}
					}
				}
			}
		}

		// execute
		err, objs := f(p, initObjs)
		if err != nil {
			return err, []interface{}{}
		}
		initObjs = objs
	}
	return nil, initObjs
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

// SkipIfNotConsensusMessage validates message type, will skip if not
func (p *Pipeline) SkipIfNotConsensusMessage(nextPhase string) *Pipeline {
	p.Items = append(p.Items, func(pipeline *Pipeline, objects ...interface{}) (error, []interface{}) {
		// check if consensus message

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
