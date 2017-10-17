package crawler

import (
	"net/http"
	"log"
	"io/ioutil"
)

func GetKuaiDaiLiTotalPage() {
	request,_ := http.NewRequest("GET","http://www.kuaidaili.com/free/",nil);
	request.Header.Set("Host","www.kuaidaili.com")
	request.Header.Set("Referer","http://www.kuaidaili.com/free/")
	request.Header.Set("DNT","1")
	request.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36")
	request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//request.Header.Set("Accept-Encoding","gzip, deflate")
	request.Header.Set("Accept-Language","zh-CN,zh;q=0.8,en;q=0.6")
	request.Header.Set("Cookie","yd_cookie=06a47780-3e46-4d0eea751414382aef2ead6e920dc3b910fc; _ydclearance=5c1877198a409475e0dee3c8-f342-43d5-b765-8a54c7c90b64-1505383772")
	client:=&http.Client{}
	resp,err:=client.Do(request)
	if err != nil {
		log.SetFlags(log.Lshortfile | log.Llongfile)
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("test.html", b, 0644)
}
