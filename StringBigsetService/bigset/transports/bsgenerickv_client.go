package transports

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	thriftpool "github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
	"sync"
)

type PoolConfig struct {
	MaxConn     uint32
	ConnTimeout uint32
	IdleTimeout uint32
}

var (
	onceInit = &sync.Once{}
	bsGenericMapPool *thriftpool.MapPool
	//bsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
	//	thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTStringBigSetKVServiceClient(c) }),
	//	thriftpool.DefaultClose)

	ibsGenericMapPool *thriftpool.MapPool
	//ibsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
	//	thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTIBSDataServiceClient(c) }),
	//	thriftpool.DefaultClose)
)

func InitPool(config PoolConfig)  {
	onceInit.Do(func() {
		bsGenericMapPool = thriftpool.NewMapPool(config.MaxConn, config.ConnTimeout, config.IdleTimeout,
			thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTStringBigSetKVServiceClient(c) }),
			thriftpool.DefaultClose)

		ibsGenericMapPool = thriftpool.NewMapPool(config.MaxConn, config.ConnTimeout, config.IdleTimeout,
			thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return generic.NewTIBSDataServiceClient(c) }),
			thriftpool.DefaultClose)
	})
}

//GetBsGenericClient client by host:port
func GetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	initBsStringIfNeed()

	client, _ := bsGenericMapPool.Get(aHost, aPort).Get()
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