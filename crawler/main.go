package crawler

import (
	"github.com/huanyusun/ProxyPool/common"
	"log"
	"github.com/PuerkitoBio/goquery"
	"os"
	"time"
	"sync"
	"github.com/huanyusun/ProxyPool/database"
	"gopkg.in/mgo.v2/bson"
)

var wg sync.WaitGroup

func Run(conf common.Conf) {
	if len(conf.Source) == 0 {
		log.Printf("Crawler Conf is Empty:\n%v", conf)
		os.Exit(2)
	}
	for _, v := range conf.Source {
		switch v {
		case "66ip":
			wg.Add(1)
			go AddProxyFrom66Ip(&wg, conf.Db.DbName)
		default:
			log.Printf("Crawler source is Error:%s", v)
			os.Exit(2)
		}

	}
	wg.Wait()
}

func AddProxyFrom66Ip(wg *sync.WaitGroup, dbName string) {
	total := GetTotalPage()
	g1 := database.GetSession()
	defer g1.Close()
	pCollection := database.GetCollection(dbName, common.GetProxysName(), g1)
	for i := 1; i <= total; i++ {
		url := SpliceUrl(i)
		query, err := goquery.NewDocument(url)
		if err != nil {
			log.Printf("Download Url is Error:%s", url)
		}
		list := GetProxyDataList(query.Selection)
		for _, v := range list {
			m := bson.M{"ip": v.Ip}
			if !database.CheckDoucumentExist(pCollection, m) {
				v.CreateTime = common.GetTimeStampNow()
				v.UpdateTime = common.GetTimeStampNow()
				pCollection.Insert(v)
			}
		}
		time.Sleep(time.Second)
	}

	wg.Done()
}
