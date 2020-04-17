package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	switch len(stages) {
	case 1:
		return stages[0](in)
	default:
		return ExecutePipeline(stages[0](in), done, stages[1:]...)
	}
}
