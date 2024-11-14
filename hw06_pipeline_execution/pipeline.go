package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	nextIn := in
	var nextOut Bi

	for _, st := range stages {
		nextOut = make(Bi)

		go func(innerIn In, innerOut Bi) {
			defer close(innerOut)
			out := st(innerIn)
			for {
				select {
				case <-done:
					return
				default:
					select {
					case <-done:
						return
					case el, ok := <-out:
						if !ok {
							return
						}
						innerOut <- el
					}
				}
			}
		}(nextIn, nextOut)

		nextIn = nextOut
	}

	return nextOut
}
