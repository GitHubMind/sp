package scheduler

import (
	"log"
	"spider/bin/engine"
)

//
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	//赋植
	s.WorkerChan = c
	//初始化一个chan
}
func (s *SimpleScheduler) CreateWorkerChan() chan engine.Request {
	//这次要不要go呢
	return make(chan engine.Request)
}
func (s *SimpleScheduler) Submit(request engine.Request) {
	//加了go 就打扒了接收和等待了
	//go func() {
	log.Println("上班")
	s.WorkerChan <- request
	//}()

}
func (s *SimpleScheduler) Clear(request engine.Request) {
	log.Println("下班")
	s.WorkerChan <- request
}
