package pipeline

// Run will execute the pipeline with the provided init objects
func Run(p *Pipeline, startingPhase string, initObjs []interface{}) (error, []interface{}) {
	// find phase to start with
	i, err := p.IndexForPhase(startingPhase)
	if err != nil {
		return err, nil
	}

	// execution loop
	for {
		if i >= len(p.Items) {
			break
		}

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
					skipToIndex, err := p.IndexForPhase(initObjs[1].(string))
					if err != nil {
						return err, nil
					}
					i = skipToIndex // bump i to phase index
					initObjs = initObjs[2:]
					continue
				}
			}
		}

		// execute
		err, objs := f(p, initObjs...)
		if err != nil {
			return err, []interface{}{}
		}
		initObjs = objs
		i++
	}
	return nil, initObjs
}
