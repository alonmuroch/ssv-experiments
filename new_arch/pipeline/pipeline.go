package pipeline

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/p2p"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/ssv"
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
	Identifier p2p.Identifier
	Runner     *ssv.Runner
	Instance   *qbft.Instance
	Items      []PipelineF
	Phase      map[string]int // maps phase name to index
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

// IndexForPhase returns the index for a marked phase or error
func (p *Pipeline) IndexForPhase(phase string) (int, error) {
	if val, found := p.Phase[phase]; found {
		return val, nil
	} else {
		return 0, errors.New("phase not found")
	}
}
