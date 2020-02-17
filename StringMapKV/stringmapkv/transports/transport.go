package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/StringMapKV" //Todo: Fix this
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/OpenStars/backendclients/go/stringmapkv/thrift/gen-go/OpenStars/Common/StringMapKV"
	thriftpool "github.com/OpenStars/thriftpoolv2"
)

var (
	mStringMapKVServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (StringMapKV.NewStringMapKVServiceClient(c)) }),
		thriftpool.DefaultClose)

	mStringMapKVServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (StringMapKV.NewStringMapKVServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift StringMapKVService client ")
}

//GetStringMapKVServiceBinaryClient client by host:port
func GetStringMapKVServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mStringMapKVServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetStringMapKVServiceCompactClient get compact client by host:port
func GetStringMapKVServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mStringMapKVServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
