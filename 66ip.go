package ProxyPool

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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
func GetProxyDataList(query *goquery.Selection) ([]PoolData) {
	list := make([]PoolData, 0, 10)
	query.Find("#main").Find("tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		data := PoolData{}
		selection.Find("td").Each(func(index int, s *goquery.Selection) {
			switch index {
			case 0:
				data.Ip = s.Text()
			case 1:
				data.Port = s.Text()
			case 2:
				data.Place = s.Text()
			case 3:
				data.ProxyType = s.Text()

			}
		})
		list = append(list, data)
	})
	return list
}
