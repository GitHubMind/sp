package engine

import (
	"fmt"
	"log"
	"spider/bin/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	//增加任务队列
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		//下一个咯
		//wg.Add(1)
		body, err := Worker(r)
		if err != nil {
			log.Println("fetch 出问题了:", r.Url, "问题是:", err)
			//这个爬虫只会跳过不会重新链接 需要改进
			log.Println("重复请求", r.Url)
			continue
		} else {
			//重复请求
			requests = requests[1:]
		}
		//解析
		//返回需要爬虫的请求，并且添加的任务列表中
		log.Println(r.Url, "请求结束")
		parseResult, err := r.ParserFun(body, r.Pre)
		requests = append(requests, parseResult.Requests...)
		//for _, item := range parseResult.Items {
		//	log.Println("Got item %s", item)
		//}

		for key, request := range requests {
			fmt.Println(key+1, ":", request)
		}

	}
	//wg.Wait()
	log.Println("finish")
}

//如果用r可能会比较好一点把对于扩展来说
func Worker(r Request) (body []byte, err error) {
	body, err = fetcher.Fetch(r.Url)
	return
}
