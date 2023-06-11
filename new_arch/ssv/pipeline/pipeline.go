package pipeline

import (
	"github.com/pkg/errors"
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

const (
	PreConsensusPhase  = "PreConsensusPhase"
	ConsensusPhase     = "ConsensusPhase"
	PostConsensusPhase = "PostConsensusPhase"
	EndPhase           = "EndPhase"
)

// PipelineF is a function taking a runner and multiple objects to process a single action
// Those functions suppose to be chained together to produce a whole working message process for a runner
type PipelineF func(runner *ssv.Runner, objects ...interface{}) (error, []interface{})

type Pipeline struct {
	Items []PipelineF
	Phase map[string]int // maps phase name to index
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Items: []PipelineF{},
		Phase: map[string]int{},
	}
}

func (p *Pipeline) Run(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	initObjs := objects
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
		err, objs := f(runner, initObjs)
		if err != nil {
			return err, []interface{}{}
		}
		initObjs = objs
	}
	return nil, initObjs
}

func (p *Pipeline) Add(f PipelineF) *Pipeline {
	p.Items = append(p.Items, f)
	return p
}

func (p *Pipeline) MarkPhase(name string) *Pipeline {
	p.Phase[name] = len(p.Items)
	return p
}

func (p *Pipeline) StopINoPreConsensusQuorum() *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		// TODO check pre-consensus quorum
		if false {
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

func (p *Pipeline) StopIfNotDecided() *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		if !runner.GetQBFT().Decided() {
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

// StopIfNoPartialSigQuorum checks if msg container for type has quorum, if not stop
func (p *Pipeline) StopIfNoPartialSigQuorum(t types.PartialSigMsgType) *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		// get container by type from runner
		// check quorum
		if false { // no quorum
			return nil, []interface{}{Stop}
		}
		return nil, objects
	})
	return p
}

func (p *Pipeline) SkipIfNotConsensusMessage(nextPhase string) *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
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

func (p *Pipeline) SkipIfNotPreConsensusMessage(nextPhase string) *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
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

func (p *Pipeline) SkipIfNotPostConsensusMessage(nextPhase string) *Pipeline {
	p.Items = append(p.Items, func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		// check if post consensus message

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
