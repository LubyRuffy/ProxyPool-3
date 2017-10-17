package main

import (
	"fmt"
	"time"
	"net/http"
	"context"
	"github.com/huanyusun/ProxyPool/common"
	"github.com/huanyusun/ProxyPool/crawler"
)

var runStatus bool

func main() {
	conf := common.GetConf()
	fmt.Printf("%v\n", conf.Db)
	ticker := time.NewTicker(time.Second * time.Duration(conf.GapsTime))
	ctx := context.Background()

	go func() {
		c, cancel := context.WithCancel(ctx)
		run(conf)
		for range ticker.C {
			runStatus = checkCtxStatus(ctx, &c, &cancel)
			if runStatus {
				fmt.Printf("[%s] Process is running, this time is ingore!\n", time.Now().Format(common.TimeFormat()))
				run(conf)
			} else {
				fmt.Printf("[%s] Process is sleep, Start now!\n", time.Now().Format(common.TimeFormat()))
				run(conf)
			}
		}
	}()

	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8090", nil)
}

func sayHello(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello test")
}

func checkCtxStatus(parent context.Context, ctx *context.Context, cancelFunc *context.CancelFunc) (status bool) {
	select {
	case <-(*ctx).Done():
		status = false
	default:
		status = true
	}
	if !status {
		*ctx, *cancelFunc = context.WithCancel(parent)
	}

	return status
}

func run(conf common.Conf) {
	crawler.Run(conf)
}
