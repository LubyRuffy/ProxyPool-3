package ProxyPool

type PoolData struct {
	Ip         string
	Port       string
	Place      string
	ProxyType  string
	CreateTime int64
	Forbidden  []string
}
