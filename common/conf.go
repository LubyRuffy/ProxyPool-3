package common

import (
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"time"
)

type Conf struct {
	Source   []string  `yaml:"source"`    //爬取数据源
	Web      WebServer `yaml:"web"`       //web服务 相关配置
	Check    []string  `yaml:"checkList"` //IP检测网址
	Db       Db        `yaml:"database"`  //数据库相关信息
	Location string    `yaml:"location"`  //时区
}

type WebServer struct {
	Open     bool   `yaml:"open"`  //是否开启
	Port     string `yaml:"port"`  //监听端口
	Https    bool   `yaml:"https"` //是否开启https
	CertFile string `yaml:"cert"`  //cert文件位置
	KeyFile  string `yaml:"key"`   //key文件位置
}

type Db struct {
	DbType   string `yaml:"dbType"`   //数据库类型
	Host     string `yaml:"host"`     //数据库host
	Port     string `yaml:"port"`     //数据库port
	Login    bool   `yaml:"login"`    //是否需要验证
	User     string `yaml:"user"`     //用户名
	PassWord string `yaml:"password"` //密码
	DbName   string `yaml:"dbName"`   //数据库名
}

var c Conf
var location *time.Location

//获取配置文件路径
func getConfFile() string {
	//return "conf/conf.yaml"
	//TODO
	return "D:/mygo/src/github.com/huanyusun/ProxyPool/cmd/conf/conf.yaml"
}

//读取配置
func (c *Conf) getConf() *Conf {
	file, err := ioutil.ReadFile(getConfFile())
	if err != nil {
		log.SetFlags(log.Llongfile | log.Lshortfile)
		log.Fatal(err)
		os.Exit(2)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.SetFlags(log.Llongfile | log.Lshortfile)
		log.Fatal(err)
		os.Exit(2)
	}
	return c
}

func init() {
	c.getConf()
	var err error
	location, err = time.LoadLocation(c.Location)
	if err != nil {
		log.Printf("Time Location is Error:\n%s", c.Location)
		os.Exit(2)
	}

}

//获取配置信息
func GetConf() Conf {
	return c
}

//获取配置下当前的时间戳
func GetTimeStampNow() int64 {
	return time.Now().In(location).Unix()
}

//获取配置下当前的时间
func GetTimeNow() string {
	return time.Now().In(location).Format("2006/01/02 15:04:05")
}

func GetProxysName() string {
	return "ProxyPool"
}

func GetExistsName() string {
	return "OverIps"
}
