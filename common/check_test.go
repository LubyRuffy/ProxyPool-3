package common

import (
	"testing"
	"github.com/huanyusun/ProxyPool"
)

func TestCheckProxy(t *testing.T) {
	pool := ProxyPool.PoolData{}
	pool.Ip = "183.232.65.202"
	pool.Port = "3128"
	pool.HttpType = false
	CheckProxy(&pool)
}
