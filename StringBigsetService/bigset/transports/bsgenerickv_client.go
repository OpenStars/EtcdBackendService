package transports

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	thriftpool "github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	bsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (generic.NewTStringBigSetKVServiceClient(c)) }),
		thriftpool.DefaultClose)

	ibsGenericMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (generic.NewTIBSDataServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetBsGenericClient client by host:port
func GetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := bsGenericMapPool.Get(aHost, aPort).Get()
	return client
}
func NewGetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := bsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}

//GetIBsGenericClient client by host:port
func GetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := ibsGenericMapPool.Get(aHost, aPort).Get()
	return client
}

func NewGetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := ibsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}
