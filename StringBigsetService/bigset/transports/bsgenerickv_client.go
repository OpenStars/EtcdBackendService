package transports

import (
	"fmt"
	"sync"

	"github.com/Sonek-HoangBui/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	thriftpool "github.com/Sonek-HoangBui/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

type PoolConfig struct {
	MaxConn     uint32
	ConnTimeout uint32
	IdleTimeout uint32
}

var (
	onceInit         = &sync.Once{}
	bsGenericMapPool *thriftpool.MapPool
	//bsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
	//	thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTStringBigSetKVServiceClient(c) }),
	//	thriftpool.DefaultClose)

	ibsGenericMapPool *thriftpool.MapPool
	mapConfig         = map[string]*PoolConfig{}
	defaultConfig     PoolConfig
	//ibsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
	//	thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTIBSDataServiceClient(c) }),
	//	thriftpool.DefaultClose)
)

func InitPool(config PoolConfig, mapConfigInit map[string]*PoolConfig) {
	onceInit.Do(func() {
		defaultConfig = config
		bsGenericMapPool = thriftpool.NewMapPool(defaultConfig.MaxConn, defaultConfig.ConnTimeout, defaultConfig.IdleTimeout,
			thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTStringBigSetKVServiceClient(c) }),
			thriftpool.DefaultClose)

		ibsGenericMapPool = thriftpool.NewMapPool(defaultConfig.MaxConn, defaultConfig.ConnTimeout, defaultConfig.IdleTimeout,
			thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTIBSDataServiceClient(c) }),
			thriftpool.DefaultClose)

		mapConfig = mapConfigInit
	})
}

//GetBsGenericClient client by host:port
func GetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	initBsStringIfNeed()
	config := mapConfig[fmt.Sprintf("%s:%s", aHost, aPort)]
	if config == nil {
		client, _ := bsGenericMapPool.GetWithConfig(aHost, aPort, defaultConfig.MaxConn, defaultConfig.ConnTimeout, 3600).Get()

		return client
	}

	client, _ := bsGenericMapPool.GetWithConfig(aHost, aPort, config.MaxConn, config.ConnTimeout, config.IdleTimeout).Get()
	return client
}

func NewGetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	initBsStringIfNeed()

	client, _ := bsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}

//GetIBsGenericClient client by host:port
func GetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	initBsStringIfNeed()

	client, _ := ibsGenericMapPool.Get(aHost, aPort).Get()
	return client
}

func NewGetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	initBsStringIfNeed()

	client, _ := ibsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}

func Close(host, port string) {
	bsGenericMapPool.Release(host, port)
}

func initBsStringIfNeed() {
	if bsGenericMapPool == nil {
		onceInit.Do(func() {
			bsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
				thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTStringBigSetKVServiceClient(c) }),
				thriftpool.DefaultClose)

			ibsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
				thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTIBSDataServiceClient(c) }),
				thriftpool.DefaultClose)
		})
	}
}
