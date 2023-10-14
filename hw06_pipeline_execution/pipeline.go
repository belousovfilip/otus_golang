package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out Bi
	for _, stage := range stages {
		out = make(Bi)
		go func(stage Stage, in In, out Bi) {
			defer close(out)
			for v := range stage(in) {
				select {
				case <-done:
					return
				default:
					out <- v
				}
			}
		}(stage, in, out)
		in = out
	}
	return out
}
