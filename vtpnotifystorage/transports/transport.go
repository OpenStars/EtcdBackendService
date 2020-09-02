package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpnotifystorage/thrift/gen-go/OpenStars/notifystorage"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mNotifyStorageVTPServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (notifystorage.NewTNotifyStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mNotifyStorageVTPServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (notifystorage.NewTNotifyStorageServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TUserCustomers client ")
}

//GetNotifyStorageVTPServiceBinaryClient client by host:port
func GetNotifyStorageVTPServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mNotifyStorageVTPServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("TUserCustomers err", err)
	}
	return client
}

//GetNotifyStorageVTPServiceCompactClient get compact client by host:port
func GetNotifyStorageVTPServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mNotifyStorageVTPServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
