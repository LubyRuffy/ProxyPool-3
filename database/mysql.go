package database

import (
	"database/sql"
	"github.com/huanyusun/ProxyPool/common"
	_"github.com/go-sql-driver/mysql"
)

var (
	dbSource string
	conf     common.Conf
	db       *sql.DB
)

//获取数据库连接
func Connect(conf common.Conf) (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = sql.Open(conf.Db.DbType, dbSource)
	}
	return db, err
}

func CheckTable() {
}

func init() {
	getDbSource()
}

func getDbSource() {
	if dbSource == "" {
		conf = common.GetConf()
		dbSource = conf.Db.User + ":" + conf.Db.PassWord + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.DbName
	}
}

func CheckDatabaseExist(db *sql.DB, dbName string) (exist bool, err error) {
	row, err := db.Query(`SHOW DATABASES LIKE '` + dbName + `'`)
	if err != nil {
		return
	}
	list, err := row.Columns()
	if err != nil {
		return
	}
	if len(list) > 0 {
		exist = true
	}
	return
}

func CreateDatabase(dbName string) (err error) {
	source := conf.Db.User + ":" + conf.Db.PassWord + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/mysql"
	db, err := sql.Open("mysql", source)
	_, err = db.Exec(`CREATE DATABASE ` + dbName)
	return
}
