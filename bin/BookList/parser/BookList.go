package parser

import (
	"fmt"
	"log"
	"regexp"
	"spider/bin/BookContent/parser"
	"spider/bin/engine"
	mongodb "spider/bin/mongdb"
)

func BookList(str []byte, pre string) (result engine.ParseResult, err error) {
	//<a href="/n/yishixiejun/7518.html"><span>第二章 君莫邪</span></a>
	res, err := regexp.Compile(`<a href="(\S+)"><span>([\p{Han} \S^]{0,50})</span></a>`)
	if err != nil {
		//log.Println(err)
	}
	Match := res.FindAllStringSubmatch(string(str), -1)
	//这是小说集合
	mongo, err := mongodb.GetSingleInstanceMongoDB("book")
	if err != nil {
		return
	}
	for key, item := range Match {
		Url := engine.HostUrl + item[1]
		//查看小说链接
		value := struct {
			Chapter_name  string
			Chapter_order string
			Book_id       string
		}{item[2], fmt.Sprint(key + 1), pre}

		pre, _ := mongo.Insert(value, "book_list")
		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookContent, Pre: pre})
	}

	return
}

func BookListPen(str []byte, pre string, mongo mongodb.Mongodb) (result engine.ParseResult, err error) {
	//<a href="/n/yishixiejun/7518.html"><span>第二章 君莫邪</span></a>
	res, err := regexp.Compile(`<h1>([\S ]+)</h1>`)
	if err != nil {
		//log.Println(err)
	}
	Matchh := res.FindStringSubmatch(string(str))
	res, err = regexp.Compile(`<div id="content">([\S \n\r]+)</div>`)
	if err != nil {
		//log.Println(err)
	}
	Match := res.FindStringSubmatch(string(str))
	//这是小说集合
	log.Println(len(Matchh))
	log.Println(len(Match))
	if len(Matchh) == 0 || len(Match) == 0 {
		log.Println(str)
		err = fmt.Errorf("返回空东西")
		return
	}
	//查看小说链接
	//log.Println(Matchh)
	//log.Println(Match)
	value := struct {
		BookName    string
		BookContent string
		Book_id     string
	}{Matchh[1], Match[1], pre}

	pre, err = mongo.Insert(value, "book_book_pen")

	//result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookContent, Pre: pre, Db: mongo})
	return
}
