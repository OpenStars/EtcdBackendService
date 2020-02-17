package transports

import (
	"log"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/OpenStars/backendclients/go/s2skv/thrift/gen-go/OpenStars/Common/S2SKV"
	thriftpool "github.com/OpenStars/thriftpoolv2"
)

var (
	s2sBinMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (S2SKV.NewTString2StringServiceClient(c)) }),
		thriftpool.DefaultClose)

	s2sCompactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (S2SKV.NewTString2StringServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	log.Println("init thrift string2stringkv client")
}

func GetS2SBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := s2sBinMapPool.Get(aHost, aPort).Get()
	return client
}

func GetS2SCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := s2sCompactMapPool.Get(aHost, aPort).Get()
	return client
}
