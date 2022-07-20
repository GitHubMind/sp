package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var str string
var wg sync.WaitGroup

const hostUrl = "http://www.quanben5.com/category/1.html"

var Book struct {
	name        string //名字
	sum         int    //章节
	isTerminate bool   //是否完结
	description string //
}

func main() {
	//fmt.Println(str)

	wg.Add(1)
	text, err := testUrl(hostUrl)
	if err != nil {
		log.Fatalln(err)
		log.Println(text)
	}
	str = string(text)
	//log.Println(str)
	cataText := catalog(text)
	//log.Println(cataText[0])
	fmt.Printf("%v", cataText[0])
	for i, s := range cataText[0] {
		fmt.Println(i, s)

	}
	wg.Wait()

}
func testUrl(url string) (text []byte, err error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		//log.Fatalln(err)
		return
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Cookie", "_ga=GA1.2.240005272.1649673036; _gid=GA1.2.1514667220.1649673036; _gat=1")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Mobile Safari/537.36")
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	} else {

		defer resp.Body.Close()
	}
	defer wg.Done()
	if resp.StatusCode == http.StatusOK {
		//e := determineEncoding(resp.Body)
		//utf8Reader :=transform.NewReader(resp.Body,e.NewEncoder())
		all, err := ioutil.ReadAll(resp.Body)
		//fmt.Printf("%s\n", all)
		text = all
		if err != nil {
			log.Println(err)

		}
	} else {
		fmt.Println("Error:status", resp.StatusCode)
		//这样可以直接结束吗
		//return fmt.Errorf("%s",)
	}

	return
}

func catalogName() {

}
func catalog(str []byte) [][]string {
	res, err := regexp.Compile(`<h3><a href="(\/n\/[a-zA-z]+\/)"><span>(\S+)</span></a></h3>([.]+|[\s]+)<p class="info">作者: <span class="author">([\p{Han}]+)</span></p>([.]+|[\s]+)<p class="description"><span class="read_ico">简</span>(\S+)</p>`)
	if err != nil {
		log.Println(err)
	}
	Match := res.FindAllStringSubmatch(string(str), -1)
	//res, err = regexp.Compile(`<p id="page_next" class="page_next"><a href="(\/category\/[0-9]+_[0-9]+.html)" rel="next">下一页</a></p>`)
	//if err != nil {
	//	log.Println(err)
	//}
	//Match = res.FindAllStringSubmatch(string(str), -1)

	return Match
}
