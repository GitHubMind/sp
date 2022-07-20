package main

import (
	"spider/bin/catalog/parser"
	"spider/bin/engine"
	mongoContorl "spider/bin/mongdb"
	"spider/bin/scheduler"
)

func main() {
	//htts: //www.qbiqu.com/wanben/1_1
	mongo, _ := mongoContorl.GetSingleInstanceMongoDB("book")

	defer mongo.Close()
	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{"https://www.qbiqu.com/wanben/1_1", parser.CatalogPen, "0", mongo})

	//engine.SimpleEngine{}.Run(engine.Request{"http://www.quanben5.com", parser.Catalog})
	//第二页测试
	//engine.Run(engine.Request{"http://www.quanben5.com/category/1.html", parser.BookCatalog, ""})
	//书本的目录
	//engine.Run(engine.Request{"http://www.quanben5.com/n/yishixiejun/xiaoshuo.html", parser.BookList, ""})
	//书本的内容
	//engine.Run(engine.Request{"http://www.quanben5.com/n/yishixiejun/xiaoshuo.html", parser.BookList})
	//<h1 class="title1">(\S+).(\S+)</h1>
	//mongo := mongodb.Connect("book_list")

	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkCount: 1}
	e.RunQuene(engine.Request{"http://www.quanben5.com", parser.Catalog, "0"})
	//e.RunQuene(engine.Request{"http://www.quanben5.com/n/yishixiejun/xiaoshuo.html", parser.BookList, "0", mongo})
	//e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkCount: 1}
	//e.RunQuene(engine.Request{"http://www.bigear.cn/reslist-2370-1.html", parser.CatalogEnglish, "0", mongo})

}
