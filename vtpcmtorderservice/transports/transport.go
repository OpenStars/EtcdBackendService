package transports

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpcmtorderservice/thrift/gen-go/OpenStars/VTPComment"
	"github.com/OpenStars/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mCommentOrderVTPServiceBinaryMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (VTPComment.NewTVTPCommentServiceClient(c)) }),
		thriftpoolv2.DefaultClose)

	mCommentOrderVTPServiceCommpactMapPool = thriftpoolv2.NewMapPool(1000, 3600, 3600,
		thriftpoolv2.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (VTPComment.NewTVTPCommentServiceClient(c)) }),
		thriftpoolv2.DefaultClose)
)

func init() {
	log.Println("init thrift TUserCustomers client ")
}

//GetCommentOrderVTPServiceBinaryClient client by host:port
func GetCommentOrderVTPServiceBinaryClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, err := mCommentOrderVTPServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("TUserCustomers err", err)
	}
	return client
}

//GetCommentOrderVTPServiceCompactClient get compact client by host:port
func GetCommentOrderVTPServiceCompactClient(aHost, aPort string) *thriftpoolv2.ThriftSocketClient {
	client, _ := mCommentOrderVTPServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
