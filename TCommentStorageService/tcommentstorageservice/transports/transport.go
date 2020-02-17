package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/TCommentStorageService" //Todo: Fix this

	"fmt"

	"github.com/OpenStars/backendclients/go/tcommentstorageservice/thrift/gen-go/OpenStars/Common/TCommentStorageService"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTCommentStorageServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} {
			return (TCommentStorageService.NewTCommentStorageServiceClient(c))
		}),
		thriftpoolv2.DefaultClose)

	mTCommentStorageServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} {
			return (TCommentStorageService.NewTCommentStorageServiceClient(c))
		}),
		thriftpoolv2.DefaultClose)
)

func init() {
	fmt.Println("init thrift TCommentStorageService client ")
}

//GetTCommentStorageServiceBinaryClient client by host:port
func GetTCommentStorageServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTCommentStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTCommentStorageServiceCompactClient get compact client by host:port
func GetTCommentStorageServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mTCommentStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
