package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTPostStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (TPostStorageService.NewTPostStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mTPostStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (TPostStorageService.NewTPostStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TPostStorageService client ")
}

//GetTPostStorageServiceBinaryClient client by host:port
func GetTPostStorageServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTPostStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTPostStorageServiceCompactClient get compact client by host:port
func GetTPostStorageServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTPostStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
