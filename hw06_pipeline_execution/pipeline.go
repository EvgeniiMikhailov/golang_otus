package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func stageWrapper(in In, done In, stage Stage) Out {
	closableIn := make(Bi)

	go func(closableIn Bi, oldIn In) {
		defer close(closableIn)

		for {
			select {
			case <-done:
				return
			case v, ok := <-oldIn:
				if !ok {
					return
				}
				closableIn <- v
			}
		}
	}(closableIn, in)
	return stage(closableIn)
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	switch len(stages) {
	case 1:
		return stageWrapper(in, done, stages[0])
	default:
		return ExecutePipeline(stageWrapper(in, done, stages[0]), done, stages[1:]...)
	}
}
