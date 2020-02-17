package transports

import (
	"log"

	"github.com/OpenStars/backendclients/go/tpubprofileservice/thrift/gen-go/openstars/pubprofile"

	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mPubProfileServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (pubprofile.NewPubProfileServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mPubProfileServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (pubprofile.NewPubProfileServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TPubProfileService client ")
}

//GetPubProfileServiceBinaryClient client by host:port
func GetPubProfileServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mPubProfileServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("GetPubProfileServiceBinaryClient err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetPubProfileServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mPubProfileServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
