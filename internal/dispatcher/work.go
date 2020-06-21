package dispatcher

import "sync"

func (dispatch *Dispatcher) work(
	wg *sync.WaitGroup,
	pipelineChannel chan Pipeliner,
	successChan chan PipelineSuccess,
	errorChannel chan PipelineError,
) {
	defer wg.Done()
	defer close(successChan)
	defer close(errorChannel)

	for {
		pipeline, more := <-pipelineChannel

		if more {
			dispatch.debugLogger.Printf("Starting pipeline: %s", pipeline.Name())
			success := pipeline.Run(errorChannel)

			if success != nil {
				successChan <- *success
			}
		} else {
			dispatch.debugLogger.Print("All pipelines complete")
			return
		}
	}

}
