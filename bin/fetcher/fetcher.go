package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func testUrl(url string, group sync.WaitGroup) (text []byte, err error) {
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
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("OK")
	defer resp.Body.Close()
	defer group.Done()
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
		//fmt.Println("Error:status", resp.StatusCode)
		//这样可以直接结束吗
		//return fmt.Errorf("%s",)
		log.Fatalln("出错了")
		err = fmt.Errorf("wrong status code :%d", resp.StatusCode)
	}
	return
}

func Fetch(urlStr string) (text []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	time.Sleep(time.Duration(num) * time.Second)
	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		//log.Fatalln(err)
		return
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Cookie", "bigear_sid=sCWnFQLzko; PHPSESSID=hgvd3lpd8uh55lmiuu1vn3le14; visitHistory=a%3A1%3A%7Bs%3A2%3A%22rs%22%3Ba%3A6%3A%7Bi%3A0%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A39%3A%22%E7%BB%8F%E5%85%B8%E5%90%8D%E4%BA%BA%E8%8B%B1%E8%AF%AD%E6%BC%94%E8%AE%B2%E9%9F%B3%E9%A2%91%E5%92%8C%E6%96%87...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700140059%22%3B%7Di%3A1%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A39%3A%22%E7%BB%8F%E5%85%B8%E5%90%8D%E4%BA%BA%E8%8B%B1%E8%AF%AD%E6%BC%94%E8%AE%B2%E9%9F%B3%E9%A2%91%E5%92%8C%E6%96%87...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700140059%22%3B%7Di%3A2%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A31%3A%22%E5%B8%83%E8%8E%B1%E5%B0%94%E9%A6%96%E7%9B%B8%E6%BC%94%E8%AE%B2%EF%BC%9AThe%20...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700143553%22%3B%7Di%3A3%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A39%3A%22%E7%BB%8F%E5%85%B8%E5%90%8D%E4%BA%BA%E8%8B%B1%E8%AF%AD%E6%BC%94%E8%AE%B2%E9%9F%B3%E9%A2%91%E5%92%8C%E6%96%87...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700140068%22%3B%7Di%3A4%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A39%3A%22%E7%BB%8F%E5%85%B8%E5%90%8D%E4%BA%BA%E8%8B%B1%E8%AF%AD%E6%BC%94%E8%AE%B2%E9%9F%B3%E9%A2%91%E5%92%8C%E6%96%87...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700140059%22%3B%7Di%3A5%3Ba%3A2%3A%7Bs%3A5%3A%22title%22%3Bs%3A39%3A%22%E7%BB%8F%E5%85%B8%E5%90%8D%E4%BA%BA%E8%8B%B1%E8%AF%AD%E6%BC%94%E8%AE%B2%E9%9F%B3%E9%A2%91%E5%92%8C%E6%96%87...%22%3Bs%3A3%3A%22url%22%3Bs%3A22%3A%22res-2370-7777700140066%22%3B%7D%7D%7D")
	//req.Header.Set("Range", "bytes=0-")
	req.Header.Set("Referer", "http://www.bigear.cn/res-2370-7777700140059.html")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Mobile Safari/537.36")

	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	//req.Header.Set("Cache-Control", "max-age=0")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Cookie", "security_session_verify=68edb851b216714bb3e144a0b71c199c; fikker-IDgO-CB34=oVNdeDkFKVSlm2SNIjilLNLhwL90Nhgr; fikker-IDgO-CB34=oVNdeDkFKVSlm2SNIjilLNLhwL90Nhgr; fikker-XaXo-BmJs=8aRLfeB909TcI9Km4UrWJu2uL98Gb9D3; fikker-XaXo-BmJs=8aRLfeB909TcI9Km4UrWJu2uL98Gb9D3")
	//req.Header.Set("If-None-Match", "1650700019|")
	//req.Header.Set("Sec-Fetch-Dest", "document")
	//req.Header.Set("Sec-Fetch-Mode", "navigate")
	//req.Header.Set("Sec-Fetch-Site", "none")
	//req.Header.Set("Sec-Fetch-User", "?1")
	//req.Header.Set("Upgrade-Insecure-Requests", "1")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Mobile Safari/537.36")
	//req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	//req.Header.Set("sec-ch-ua-mobile", "?1")
	//req.Header.Set("sec-ch-ua-platform", `"Android"`)
	//proxyAddr := "http://183.247.215.218:30001/"
	//proxy, err := url.Parse(proxyAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//netTransport := &http.Transport{
	//	Proxy:                 http.ProxyURL(proxy),
	//	MaxIdleConnsPerHost:   10,
	//	ResponseHeaderTimeout: time.Second * time.Duration(5),
	//}

	client := &http.Client{
		//Transport: netTransport,
		Timeout: 4 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("过时操作", err)
		//保险一点嘛
		return
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			//e := determineEncoding(resp.Body)
			r := bufio.NewReader(resp.Body)
			log.Println(r.Size())
			if r.Size() == 0 {
				err = fmt.Errorf("请求放回的数据为0")
				return
			}
			utf8Reader, _ := charset.NewReader(resp.Body, "UTF-8	")
			text, err = ioutil.ReadAll(utf8Reader)
			if err != nil {
				//log.Println("StatusOK:", err)
				err = fmt.Errorf("StatusOK:%s", err)
				return
			}
		} else {
			//fmt.Println("Error:status", resp.StatusCode)
			//这样可以直接结束吗
			//return fmt.Errorf("%s",)
			//log.Fatalln("出错了")
			err = fmt.Errorf("wrong status code :%d", resp.StatusCode)
		}
	}
	//group.Done()
	return
}
