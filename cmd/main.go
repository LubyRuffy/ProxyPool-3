package main

import (
	"github.com/huanyusun/ProxyPool/common"
	"fmt"
	"github.com/huanyusun/ProxyPool/crawler"
)

func main() {
	c := common.GetConf()
	fmt.Printf("%v", c.Db)
	crawler.Run(c)
}
