package scheduler

import "spider/bin/engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(w chan engine.Request) {

}

func (s *QueuedScheduler) CreateWorkerChan() ( chan engine.Request) {
	//这次要不要go呢
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	//这次要不要go呢
	s.RequestChan <- r
}

func (s *QueuedScheduler) WorderReadey(w chan engine.Request) {
	s.WorkerChan <- w
}
func (s *QueuedScheduler) Run() {
	//这样定义一个新的就可以了
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//活动过
			if len(requestQ) > 0 && len(workerQ) > 0 {
				//如果有就去worker
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

//func ConfigureMasterWorkerChan(chan)
