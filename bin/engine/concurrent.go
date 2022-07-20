package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(request Request)
	//测试queune 注释的
	ConfigureMasterWorkerChan(chan Request)
	WorderReadey(w chan Request)
	CreateWorkerChan() chan Request
	Run()
}

func (e *ConcurrentEngine) RunQuene(seeds ...Request) {
	e.Scheduler.Run()
	out := make(chan ParseResult)
	//为何要两个
	for i := 0; i < e.WorkCount; i++ {
		createWorkeQuqnue(e.Scheduler.CreateWorkerChan(), out, e.Scheduler)
	}
	for _, item := range seeds {
		e.Scheduler.Submit(item)
	}
	for {
		//开始并发处理
		result := <-out
		//为什么要打赢item,
		for key, request := range result.Requests {
			//送入任务列表
			log.Println(key, "送入任务", request.Url)
			//这个go了之后不会堵塞了
			e.Scheduler.Submit(request)

		}
	}
}
func createWorkeQuqnue(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			//log.Println("下班!!!")
			s.WorderReadey(in)
			request := <-in
			result, err := SimpleWorker(request)
			if err != nil {
				//如果出错了 怎么解决呢？
				fmt.Println("重新请求:", request.Url)
				//我想重新请求
				//in <- request
				//这个可能也要修改 因为不再是通过这来蓝判断等待了
				out <- NilParseResult(request)
				continue
			}
			log.Println(request.Url, "请求结束")
			out <- result
		}
	}()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//for _, seed := range seeds {
	//	//难道submit是任务队列的动作？
	//	e.Scheduler.Submit(  seed)
	//}
	//声明队列
	in := make(chan Request)
	out := make(chan ParseResult)
	//为何要两个
	//测试quenue注释的
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}
	for _, item := range seeds {
		e.Scheduler.Submit(item)
	}
	for {
		//开始并发处理
		result := <-out
		//为什么要打赢item,
		//for _, item := range result.Items {
		//	fmt.Println(item)
		//}
		log.Println(result.Requests)
		for key, request := range result.Requests {
			//送入任务列表
			log.Println(key, "送入任务", request.Url)
			//这个go了之后不会堵塞了
			e.Scheduler.Submit(request)

		}

	}
	log.Println(13)
}
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			//log.Println("下班!!!")
			request := <-in
			result, err := SimpleWorker(request)
			if err != nil {
				//如果出错了 怎么解决呢？
				fmt.Println("重新请求:", request.Url)
				//我想重新请求
				//in <- request
				out <- NilParseResult(request)
				continue
			}
			log.Println(request.Url, "请求结束")
			out <- result
		}
	}()
}
