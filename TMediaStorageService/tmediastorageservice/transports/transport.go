package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/TMediaStorageService" //Todo: Fix this
	"fmt"

	"github.com/OpenStars/EtcdBackendService/TMediaStorageService/tmediastorageservice/thrift/gen-go/OpenStars/Common/TMediaStorageService"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTMediaStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (TMediaStorageService.NewTMediaStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mTMediaStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (TMediaStorageService.NewTMediaStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	fmt.Println("init thrift TMediaStorageService client ")
}

//GetTMediaStorageServiceBinaryClient client by host:port
func GetTMediaStorageServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTMediaStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTMediaStorageServiceCompactClient get compact client by host:port
func GetTMediaStorageServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTMediaStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
