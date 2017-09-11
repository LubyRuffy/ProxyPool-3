package crawler

import (
	"github.com/huanyusun/ProxyPool/common"
	"log"
	"github.com/PuerkitoBio/goquery"
	"os"
	"time"
	"sync"
	"github.com/huanyusun/ProxyPool/database"
	"fmt"
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
	//g1 := database.GetSession()
	g2 := database.GetSession()
	//defer g1.Close()
	defer g2.Close()
	//pCollection := database.GetCollection(dbName, common.GetProxysName(), g1)
	rCollection := database.GetCollection(dbName, common.GetExistsName(), g2)
	//m := make([]bson.M, 0, 2)
	//m = append(m, bson.M{"ip": "127.0.0.1"})
	//m = append(m, bson.M{"ip": "192.168.0.1"})
	m := bson.M{"ip": "127.0.0.1"}
	rCollection.Insert(m)
	os.Exit(1)
	for i := 1; i <= total; i++ {
		url := SpliceUrl(i)
		query, err := goquery.NewDocument(url)
		if err != nil {
			log.Printf("Download Url is Error:%s", url)
		}
		_, ips := GetProxyDataList(query.Selection)
		fmt.Printf("%v\n", ips)

		//pCollection.Insert(list)
		//rCollection.Insert(ips)
		time.Sleep(time.Second)
	}

	wg.Done()
}

//func GetParams(list []ProxyPool.PoolData) (m bson.M) {
//	for _, v := range list {
//		buf := bytes.Buffer{}
//		buf.Write(v.Ip)
//		buf.WriteString(":")
//		buf.Write(v.Port)
//		//m["ip"] =
//	}
//}
