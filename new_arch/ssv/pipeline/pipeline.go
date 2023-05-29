package pipeline

import "ssv-experiments/new_arch/ssv"

type PipelineControlSymbols uint64

const (
	// Stop means stop pipeline
	Stop PipelineControlSymbols = iota
	// SkipNext means skip the next pipeline element
	SkipNext
)

func IsSkipNext(objs ...interface{}) bool {
	return isPipelineSymbol(SkipNext, objs)
}

func IsStop(objs ...interface{}) bool {
	return isPipelineSymbol(Stop, objs)
}

func isPipelineSymbol(s PipelineControlSymbols, objs ...interface{}) bool {
	if len(objs) > 0 {
		if c, isControlSymbol := objs[0].(PipelineControlSymbols); isControlSymbol && c == s {
			return true
		}
	}
	return false
}

// PipelineF is a function taking a runner and multiple objects to process a single action
// Those functions suppose to be chained together to produce a whole working message process for a runner
type PipelineF func(runner *ssv.Runner, objects ...interface{}) (error, []interface{})

type Pipeline struct {
	Items []PipelineF
}

func NewPipeline() *Pipeline {
	return &Pipeline{Items: []PipelineF{}}
}

func (p *Pipeline) Run(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
	initObjs := objects
	for _, f := range p.Items {
		// check control symbols
		if len(initObjs) > 0 {
			if c, isControlSymbol := initObjs[0].(PipelineControlSymbols); isControlSymbol {
				switch c {
				case Stop:
					return nil, []interface{}{}
				case SkipNext:
					continue
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

func (p *Pipeline) AddIfConsensusMessage(f PipelineF) *Pipeline {
	p.Items = append(p.Items, ContinueIfConsensusMessage(f))
	return p
}

func (p *Pipeline) AddIfPreConsensusMessage(f PipelineF) *Pipeline {
	p.Items = append(p.Items, ContinueIfPreConsensusMessage(f))
	return p
}

func (p *Pipeline) AddIfPostConsensusMessage(f PipelineF) *Pipeline {
	p.Items = append(p.Items, ContinueIfPostConsensusMessage(f))
	return p
}

// Pipepify turns a pipeline into a pipeline function to nest pipelines in pipelines
func (p *Pipeline) Pipepify() PipelineF {
	return func(runner *ssv.Runner, objects ...interface{}) (error, []interface{}) {
		return p.Run(runner, objects)
	}
}
