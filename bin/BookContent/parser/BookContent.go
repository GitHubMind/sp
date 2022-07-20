package parser

import (
	"log"
	"regexp"
	"spider/bin/engine"
	mongoContorl "spider/bin/mongdb"
)

func BookContent(str []byte, pre string) (result engine.ParseResult, err error) {
	//<a href="/n/yishixiejun/7518.html"><span>第二章 君莫邪</span></a>
	res, err := regexp.Compile(`<h1 class="title1">((\S+).(\S+))</h1>`)
	if err != nil {
		//log.Println(err)
	}
	Match := res.FindAllStringSubmatch(string(str), -1)
	var title string
	for _, item := range Match {
		title = item[1]
	}
	//fmt.Println(string(str))
	//([\S ]*)([\f\n\r\t\v])</div>
	res, err = regexp.Compile(`<div id="content">([\f\n\r\t\v]*)([\S ]*)`)
	//res, err = regexp.Compile(`[\p{Han}]+`)
	if err != nil {
		log.Println(err)
	}
	Match = res.FindAllStringSubmatch(string(str), -1)
	//这是页数
	mongo, err := mongoContorl.GetSingleInstanceMongoDB("book")
	if err != nil {
		return
	}
	for _, item := range Match {
		//fmt.Println(item[2])
		//Url := engine.HostUrl + item[1]
		//查看小说链接
		value := struct {
			Contnet      string
			Book_list_id string
			Title        string
		}{item[2], pre, title}
		mongo.Insert(value, "book_content")

	}
	return
}
