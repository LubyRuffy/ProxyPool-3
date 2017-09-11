package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"github.com/huanyusun/ProxyPool"
	"github.com/huanyusun/ProxyPool/common"
	"bytes"
)

var mainUrl string = "http://www.66ip.cn/"

//获取数据页数
func GetTotalPage() int {
	resp, err := http.Get(mainUrl)
	if err != nil {
		log.Fatal(err)
	}
	query, _ := goquery.NewDocumentFromReader(resp.Body)
	t := query.Selection.Find(".dotdot").Next().Text()

	if err != nil {
		log.Fatal(err)
	}
	total, _ := strconv.Atoi(t)
	return total
}

//拼接返回数据地址
func SpliceUrl(i int) string {
	return mainUrl + strconv.Itoa(i)
}

//获取IP数据切片
func GetProxyDataList(query *goquery.Selection) ([]ProxyPool.PoolData, []string) {
	list := make([]ProxyPool.PoolData, 0, 10)
	ipList := make([]string, 0, 10)
	query.Find("#main").Find("tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		data := ProxyPool.PoolData{}
		selection.Find("td").Each(func(index int, s *goquery.Selection) {
			text := []byte(s.Text())
			text, _ = common.GbkToUtf8(text)
			switch index {
			case 0:
				data.Ip = []byte(s.Text())
			case 1:
				data.Port = []byte(s.Text())
			case 2:
				data.Place = []byte(s.Text())
			case 3:
				data.ProxyType = []byte(s.Text())
			}
		})
		data.HttpType = 1
		buffer := bytes.Buffer{}
		buffer.Write(data.Ip)
		buffer.WriteString(":")
		buffer.Write(data.Port)
		list = append(list, data)
		ipList = append(ipList, buffer.String())
	})
	return list, ipList
}
