package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Platform/MarketPlace" //Todo: Fix this

	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/OpenStars/EtcdBackendService/MarketPlaceService/marketplaceitem/thrift/gen-go/OpenStars/Platform/MarketPlace"
	"github.com/OpenStars/thriftpool"
)

var (
	mTMarketPlaceServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (MarketPlace.NewTMarketPlaceServiceClient(c)) }),
		thriftpool.DefaultClose)

	mTMarketPlaceServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (MarketPlace.NewTMarketPlaceServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift TMarketPlaceService client ")
}

//GetTMarketPlaceServiceBinaryClient client by host:port
func GetTMarketPlaceServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMarketPlaceServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTMarketPlaceServiceCompactClient get compact client by host:port
func GetTMarketPlaceServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMarketPlaceServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
