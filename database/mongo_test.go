package database

import (
	"testing"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func TestGetSession(t *testing.T) {
	session := GetSession()
	fmt.Println(*session)
}

func TestGetCollection(t *testing.T) {
	c := GetCollection("ProxyPool", "ProxyPool")
	m := bson.M{}
	result := make(map[string]string)
	c.Find(m).Distinct("ip", &result)
	c.Insert()
	fmt.Printf("%v", result)
}
