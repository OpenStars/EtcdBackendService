package transports

import (
	"github.com/OpenStars/EtcdBackendService/signservice/gen-go/openstars/signservice"

	thriftpool "github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mSignServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (signservice.NewSignDataServiceClient(c)) }),
		thriftpool.DefaultClose)

	mSignServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (signservice.NewSignDataServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetSignServiceBinaryClient client by host:port
func GetSignServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mSignServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetSignServiceCompactClient get compact client by host:port
func GetSignServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mSignServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
