package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/vtporderinfoservice2/thrift/gen-go/OpenStars/orderservice"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (orderservice.NewTOrderServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (orderservice.NewTOrderServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TUserCustomers client ")
}

//GetOrderInfoServiceBinaryClient client by host:port
func GetOrderInfoServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("TUserCustomers err", err)
	}
	return client
}

//GetOrderInfoServiceCompactClient get compact client by host:port
func GetOrderInfoServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
