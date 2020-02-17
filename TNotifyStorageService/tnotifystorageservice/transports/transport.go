package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/TNotifyStorageService" //Todo: Fix this
	//Todo: Fix this
	"fmt"

	"github.com/OpenStars/backendclients/go/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTNotifyStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} {
			return (TNotifyStorageService.NewTNotifyStorageServiceClient(c))
		}),
		thriftpoolv2.DefaultClose)

	mTNotifyStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (TNotifyStorageService.NewTNotifyStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	fmt.Println("init thrift TNotifyStorageService client ")
}

//GetTNotifyStorageServiceBinaryClient client by host:port
func GetTNotifyStorageServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTNotifyStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTNotifyStorageServiceCompactClient get compact client by host:port
func GetTNotifyStorageServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTNotifyStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
