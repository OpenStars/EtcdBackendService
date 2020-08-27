package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpcalllogservice/thrift/gen-go/OpenStars/calllog"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mCallLogVTPServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (calllog.NewTCallLogServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mCallLogVTPServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (calllog.NewTCallLogServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TUserCustomers client ")
}

//GetCallLogVTPServiceBinaryClient client by host:port
func GetCallLogVTPServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mCallLogVTPServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("TUserCustomers err", err)
	}
	return client
}

//GetCallLogVTPServiceCompactClient get compact client by host:port
func GetCallLogVTPServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mCallLogVTPServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
