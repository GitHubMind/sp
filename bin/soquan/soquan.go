package soquan

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
