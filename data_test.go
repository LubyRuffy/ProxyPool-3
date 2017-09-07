package ProxyPool

import (
	"testing"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func TestGetSession(t *testing.T) {
	session := GetSession()
	fmt.Println(*session)
}

func TestAddProxy(t *testing.T) {
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
		AddProxy(v)
	}
}

func TestGetProxy(t *testing.T) {
	m := map[string]string{"place": "印度尼西亚"}
	data := GetProxy("ProxyPool", "ProxyPool", m)
	for k, v := range data {
		fmt.Println(k, v)
	}

}
