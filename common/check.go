package common

import (
	"net/http"
	"github.com/huanyusun/ProxyPool"
	"net/url"
	"log"
	"crypto/tls"
	"time"
	"fmt"
	"io/ioutil"
	"sync"
)

func CheckProxy(pool *ProxyPool.PoolData) {

	proxyUrl := ""
	if pool.HttpType {
		proxyUrl = "https://"
	} else {
		proxyUrl = "http://"
	}
	proxyUrl = proxyUrl + pool.Ip + ":" + pool.Port
	fmt.Println(proxyUrl)
	pr, err := url.Parse(string(proxyUrl))
	if err != nil {
		log.SetFlags(log.Llongfile | log.Lshortfile)
		log.Fatal(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(pr),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}
	c := GetConf()
	wg := sync.WaitGroup{}
	for _, v := range c.CheckList.MainLand {
		fmt.Println(v)
		wg.Add(1)
		go CheckUrl(v, client, &wg)
	}
	wg.Wait()
}

func CheckUrl(url string, client *http.Client, wg *sync.WaitGroup) {
	request, _ := http.NewRequest("GET", "http://"+url, nil)
	request.Header.Set("Pragma", "no-cache")
	//request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36")
	request.Header.Set("Accept", "image/webp,image/apng,image/*,*/*;q=0.8")
	request.Header.Set("Referer", url)
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	sr, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(url+".html", sr, 0644)
	wg.Done()
}
