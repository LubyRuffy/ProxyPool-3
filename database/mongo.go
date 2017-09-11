package database

import (
	"gopkg.in/mgo.v2"
	"github.com/huanyusun/ProxyPool"
)

const url = "127.0.0.1:27017"

var (
	goSession *mgo.Session
)

//获取连接session
func GetSession() *mgo.Session {
	if goSession == nil {
		var err error
		goSession, err = mgo.Dial(url)
		if err != nil {
			panic(err)
		}
	}
	return goSession.Clone()
}

//获取数据集
func GetCollection(dbName string, collection string, g *mgo.Session) (c *mgo.Collection) {
	db := g.DB(dbName)
	c = db.C(collection)
	return
}

//添加数据
func AddProxy(list []ProxyPool.PoolData, collection *mgo.Collection) {
	collection.Insert(ProxyPool.PoolData{})
}

func CheckDataRepeat(ipList []string, collection *mgo.Collection) []string {
	ips := make([]string, 0, 10)
	collection.Find(ipList).All(&ips)
	return ips
}
