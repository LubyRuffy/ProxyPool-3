package ProxyPool

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const url = "127.0.0.1:27017"

var (
	goSession *mgo.Session
	location  *time.Location
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

//获取时区
func GetLocation() *time.Location {
	if location == nil {
		location, _ = time.LoadLocation("Asia/Shanghai")
	}
	return location
}

//增加数据
func AddProxy(p PoolData) {
	g := GetSession()
	db := g.DB("ProxyPool")
	c := db.C("ProxyPool")
	p.CreateTime = time.Now().In(GetLocation()).Unix()
	c.Insert(p)
}

//获取数据
func GetProxy(dbName string, collection string, params map[string]string) []interface{} {
	g := GetSession()
	db := g.DB(dbName)
	c := db.C(collection)
	finds := bson.M{}
	for k, v := range params {
		finds[k] = v
	}
	data := make([]interface{}, 0, 0)
	c.Find(finds).All(&data)
	return data
}
