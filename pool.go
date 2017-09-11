package ProxyPool

type PoolData struct {
	Ip         []byte
	Port       []byte
	Place      []byte
	ProxyType  []byte
	CreateTime int64
	UpdateTime int64
	HttpType   int8
}
