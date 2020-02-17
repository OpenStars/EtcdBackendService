package transports

import (
	"log"

	"github.com/OpenStars/backendclients/go/tmediacloudservice/thrift/gen-go/openstars/mcloud"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mMediaCloudServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (mcloud.NewTMediaServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mMediaCloudServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (mcloud.NewTMediaServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TMediaCloudService client ")
}

//GetMediaCloudServiceBinaryClient client by host:port
func GetMediaCloudServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mMediaCloudServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetMediaCloudServiceCompactClient get compact client by host:port
func GetMediaCloudServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mMediaCloudServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
