package ProxyPool

import (
	"testing"
	"reflect"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"os"
	"fmt"
)

func TestGetTotalPage(t *testing.T) {
	i := GetTotalPage()
	if reflect.TypeOf(i).String() == "int" && i > 0 {
		println(t.Name() + " success")
	} else {
		t.Fail()
	}
}

func TestSpliceUrl(t *testing.T) {
	total := GetTotalPage()
	for i := 1; i <= total; i++ {
		url := SpliceUrl(i)
		resp, err := http.Get(url)
		if resp.StatusCode == 200 {
			println("Success:" + url)
		}
		if err != nil {
			println(err)
			t.FailNow()
		}
	}
}

func TestGetProxyDataList(t *testing.T) {
	url := SpliceUrl(1)
	println(url)
	query, err := goquery.NewDocument(url)
	if err != nil {
		println(err)
		t.FailNow()
		os.Exit(1)
	}
	list := GetProxyDataList(query.Selection)
	for _, v := range list {
		fmt.Printf("IP:%s,PORT:%s,PLACE:%s,TYPE:%s\n", v.Ip, v.Port, v.Place, v.ProxyType)
	}
}
