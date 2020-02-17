package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV" //Todo: Fix this
	"fmt"

	"github.com/OpenStars/backendclients/go/mapphone2pubkey/thrift/gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV" //Todo: Fix this
	thriftpool "github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mTMapPhoneNumberPubkeyKVServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} {
			return (MapPhoneNumberPubkeyKV.NewTMapPhoneNumberPubkeyKVServiceClient(c))
		}),
		thriftpool.DefaultClose)

	mTMapPhoneNumberPubkeyKVServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} {
			return (MapPhoneNumberPubkeyKV.NewTMapPhoneNumberPubkeyKVServiceClient(c))
		}),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift TMapPhoneNumberPubkeyKVService client ")
}

//GetTMapPhoneNumberPubkeyKVServiceBinaryClient client by host:port
func GetTMapPhoneNumberPubkeyKVServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMapPhoneNumberPubkeyKVServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTMapPhoneNumberPubkeyKVServiceCompactClient get compact client by host:port
func GetTMapPhoneNumberPubkeyKVServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTMapPhoneNumberPubkeyKVServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
