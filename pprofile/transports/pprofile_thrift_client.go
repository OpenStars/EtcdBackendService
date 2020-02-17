package transports


import (
	"github.com/OpenStars/backendclients/go/pprofile/thrift/gen-go/OpenStars/Platform/Profile"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	thriftpool "github.com/OpenStars/thriftpool"	
	)


var (
	ppProfileBinMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFunc( func( c thrift.TClient ) (interface{}) { return  (Profile.NewTPlatformProfileServiceClient(c)) }),
		thriftpool.DefaultClose)

	ppProfileCompactMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFuncCompactProtocol( func(c thrift.TClient) (interface{}) { return  (Profile.NewTPlatformProfileServiceClient(c)) }),
		thriftpool.DefaultClose)
		
	)

 


func init(){
	fmt.Println("init thrift kvcounter client ");
}

//GetPProfileBinaryClient client by host:port
func GetPProfileBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := ppProfileBinMapPool.Get(aHost, aPort).Get()
	return client;
}

//GetPProfileCompactClient get compact client by host:port
func GetPProfileCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := ppProfileCompactMapPool.Get(aHost, aPort).Get()
	return client;
}


