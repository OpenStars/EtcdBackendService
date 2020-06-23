package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/TReportItemService/treportitemservice/thrift/gen-go/OpenStars/Platform/MarketPlace"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTReportStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (MarketPlace.NewTReportItemServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mTReportStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (MarketPlace.NewTReportItemServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TReportStorageService client ")
}

//GetTReportStorageServiceBinaryClient client by host:port
func GetTReportItemServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTReportStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTReportStorageServiceCompactClient get compact client by host:port
func GetTReportItemServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTReportStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
