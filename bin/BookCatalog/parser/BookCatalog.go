package parser

import (
	"fmt"
	"regexp"
	"spider/bin/BookList/parser"
	"spider/bin/engine"
	mongodb "spider/bin/mongdb"
)

func BookCatalog(str []byte, pre string) (result engine.ParseResult, err error) {

	res, err := regexp.Compile(`<h3><a href="(\/n\/[a-zA-z]+\/)"><span>(\S+)</span></a></h3>([.]+|[\s]+)<p class="info">作者: <span class="author">([\p{Han}]+)</span></p>([.]+|[\s]+)<p class="description"><span class="read_ico">简</span>(\S+)</p>`)
	if err != nil {
		//log.Println(err)
	}
	Match := res.FindAllStringSubmatch(string(str), -1)
	//这是小说集合
	mongo, err := mongodb.GetSingleInstanceMongoDB("book")
	if err != nil {
		return
	}
	res, _ = regexp.Compile(`div class="pic"><img src="([\S]+)"`)
	MatchSign := res.FindStringSubmatch(string(str))
	if len(MatchSign) == 0 {
		err = fmt.Errorf("%s", "")
		return
	}
	imgUrl := MatchSign[1]
	for _, item := range Match {
		Url := engine.HostUrl + item[1] + "xiaoshuo.html"
		//查看小说链接
		value := struct {
			Name       string
			Describe   string
			ImgUrl     string
			Author     string
			Catalog_id string
		}{item[2], item[6], imgUrl, item[4], pre}
		pre, _ := mongo.Insert(value, "book")
		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookList, Pre: pre})
	}
	res, err = regexp.Compile(`<p id="page_next" class="page_next"><a href="(\/category\/[0-9]+_[0-9]+.html)" rel="next">下一页</a></p>`)
	if err != nil {
		//log.Println(err)
	}
	Match = res.FindAllStringSubmatch(string(str), -1)
	for _, item := range Match {
		Url := engine.HostUrl + item[1]
		//下一页
		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: BookCatalog, Pre: pre})
	}
	return
}

//
//func BookCatalogPen(str []byte, pre string, mongo mongodb.Mongodb) (result engine.ParseResult, err error) {
//	res, err := regexp.Compile(`<h1>([\S]+)</h1>`)
//	if err != nil {
//		//log.Println(err)
//	}
//	Match := res.FindStringSubmatch(string(str))
//	//log.Println(Match)
//	bookName := Match[1]
//
//	res, _ = regexp.Compile(`<p>作&nbsp;&nbsp;&nbsp;&nbsp;者：([\S]+)</p>`)
//	Match = res.FindStringSubmatch(string(str))
//	auther := Match[1]
//	//这是小说集合
//	res, err = regexp.Compile(`<dd><a href="([\S]+)">([\S|.]+)</a></dd>`)
//	Matchh := res.FindAllStringSubmatch(string(str), -1)
//	for key, item := range Matchh {
//		Url := engine.HostUrlPen + string(item[1])
//		//查看小说链接
//		value := struct {
//			BookName    string
//			Auther      string
//			Order       string
//			CatalogName string
//			book_id     string
//		}{bookName, auther, fmt.Sprintf("%d", key), item[2], pre}
//		pre, _ := mongo.Insert(value, "book_pen")
//		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookListPen, Pre: pre, Db: mongo})
//	}
//
//	return
//}
