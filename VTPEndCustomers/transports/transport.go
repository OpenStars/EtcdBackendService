package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/VTPEndCustomers/thrift/gen-go/openstars/enduservtp"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mEndUserVTPServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (enduservtp.NewTEndUserVTPServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mEndUserVTPServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (enduservtp.NewTEndUserVTPServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TUserCustomers client ")
}

//GetPubProfileServiceBinaryClient client by host:port
func GetEndUserVTPServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mEndUserVTPServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("TUserCustomers err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetEndUserVTPServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mEndUserVTPServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
