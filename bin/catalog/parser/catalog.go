package parser

import (
	"regexp"
	"spider/bin/BookCatalog/parser"
	"spider/bin/engine"
	mongodb "spider/bin/mongdb"
)

func Catalog(str []byte, s string) (result engine.ParseResult, err error) {

	//log.Println(string(str))

	res, _ := regexp.Compile(`(\/(category)\/([0-9])\.html)"><span>([\S]+)</span>`)
	Match := res.FindAllStringSubmatch(string(str), -1)
	result = engine.ParseResult{}
	mongo, err := mongodb.GetSingleInstanceMongoDB("book")
	if err != nil {
		return
	}
	for _, item := range Match {
		Url := engine.HostUrl + item[1]
		value := struct {
			Name string
		}{item[4]}
		pre, _ := mongo.Insert(value, "catalog")
		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookCatalog, Pre: pre})
	}
	return
}

//
//func CatalogPen(str []byte, s string, mongo mongodb.Mongodb) (result engine.ParseResult, err error) {
//
//	res, err := regexp.Compile(`<a href="([\S]+)" class="next">`)
//	if err != nil {
//		//log.Println(err)
//	}
//	url := res.FindStringSubmatch(string(str))
//	//log.Println(Match)
//	if len(url) == 0 {
//		err = fmt.Errorf("url不足")
//		return
//	}
//	Url := engine.HostUrlPen + string(url[1])
//	result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: CatalogPen, Pre: "", Db: mongo})
//
//	res, _ = regexp.Compile(`<li><span class="s2">《<a href="([\S]+)" target="_blank">([\S]+)</a>》</span><span class="s3"><a href="([\S]+)" target="_blank">([\S]+)</a>([\S]+)</span><span class="s5">([\S]+)</span></li>`)
//	Match := res.FindAllStringSubmatch(string(str), -1)
//	result = engine.ParseResult{}
//	mongo = mongodb.Connect("calalog")
//
//	for _, item := range Match {
//		Url := item[1]
//		value := struct {
//			Name       string
//			Auther     string
//			LastUpdate string
//		}{item[2], item[6], item[5]}
//		pre, _ := mongo.Insert(value, "catalog_pen")
//		result.Requests = append(result.Requests, engine.Request{Url: Url, ParserFun: parser.BookCatalogPen, Pre: pre, Db: mongo})
//	}
//	return
//}
