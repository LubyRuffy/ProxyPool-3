package database

import (
	"testing"
	"github.com/huanyusun/ProxyPool/common"
	"fmt"
	"os"
)

func TestConnect(t *testing.T) {
	conf := common.GetConf()
	db, err := Connect(conf)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", db)
	rows, err := db.Query("show tables ")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	list, err := rows.Columns()
	if err != nil {
		fmt.Printf("%s", err)
	}
	for k, v := range list {
		fmt.Printf("%s:%s", k, v)
	}

}

func TestCheckDatabaseExist(t *testing.T) {
	conf := common.GetConf()
	db, err := Connect(conf)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	exist, err := CheckDatabaseExist(db, conf.Db.DbName)
	fmt.Println(exist)
}

func TestCreateDatabase(t *testing.T) {
	conf := common.GetConf()
	db, err := Connect(conf)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	exist, err := CheckDatabaseExist(db, conf.Db.DbName)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	if !exist {
		err = CreateDatabase(conf.Db.DbName)
		if err!=nil {
			fmt.Printf("%s\n", err)
		}
	}

	defer db.Close()
}
