package ProxyPool

import "gopkg.in/mgo.v2"

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

func AddProxy(p PoolData) {
	g := GetSession()
	db := g.DB("ProxyPool")
	c := db.C("ProxyPool")
	c.Insert(p)
}
