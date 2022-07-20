package parser

import (
	"fmt"
	"log"
	"regexp"
	"spider/bin/engine"
	mongodb "spider/bin/mongdb"
)

func CatalogEnglish(str []byte, s string, mongo mongodb.Mongodb) (result engine.ParseResult, err error) {

	//res, err := regexp.Compile(`<a href="([\S]+)" class="next">`)
	//if err != nil {
	//	//log.Println(err)
	//}
	//url := res.FindStringSubmatch(string(str))
	////log.Println(Match)
	//if len(url) == 0 {
	//	err = fmt.Errorf("url不足")
	//	return
	//}
	//Url := engine.HostUrlPen + string(url[1])
	//result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: engine.NilParser([]byte, ""), Pre: "", Db: mongo})

	res, _ := regexp.Compile(`<div><strong><a href="([\S]+)"`)
	Match := res.FindAllStringSubmatch(string(str), -1)
	result = engine.ParseResult{}
	mongo = mongodb.Connect("calalog")

	for _, item := range Match {
		//log.Println(item)
		Url := "http://www.bigear.cn" + item[1]
		//Url := item[1]
		//value := struct {
		//	Name       string
		//	Auther     string
		//	LastUpdate string
		//}{item[2], item[6], item[5]}
		//pre, _ := mongo.Insert(value, "catalog_pen")
		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: EnglishConent, Pre: "", Db: mongo})
	}
	return
}

func EnglishConent(str []byte, s string, mongo mongodb.Mongodb) (result engine.ParseResult, err error) {

	res, err := regexp.Compile(`<h1 id="ArticleTit" class="contenttitlename">([\S \n\r]+)</h1>`)
	if err != nil {
		//log.Println(err)
	}
	name := res.FindStringSubmatch(string(str))
	//log.Println(Match)
	if len(name) == 0 {
		err = fmt.Errorf("url不足")
		return
	}
	res, err = regexp.Compile(`<audio style="width:300px;margin-top:20px" src="([\S]+)"`)
	if err != nil {
		//log.Println(err)
	}
	listen := res.FindStringSubmatch(string(str))
	//log.Println(Match)
	if len(name) == 0 {
		log.Println("?")
		err = fmt.Errorf("url不足")
		return
	}
	value := struct {
		Name string
		Url  string
	}{name[1], listen[1]}
	mongo.Insert(value, "english_listen_cps")
	//result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: engine.NilParser([]byte, ""), Pre: "", Db: mongo})

	return
}
