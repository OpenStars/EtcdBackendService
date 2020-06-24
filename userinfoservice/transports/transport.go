package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/userinfoservice/thrift/gen-go/openstars/userinfoservice"

	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mUserInfoServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (userinfoservice.NewTUserInfoServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mUserInfoServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (userinfoservice.NewTUserInfoServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TPubProfileService client ")
}

//GetPubProfileServiceBinaryClient client by host:port
func GetUserInfoServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mUserInfoServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("GetPubProfileServiceBinaryClient err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetUserInfoServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mUserInfoServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
