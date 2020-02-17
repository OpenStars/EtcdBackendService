package transports


import (
	"github.com/OpenStars/backendclients/go/ppassport/thrift/gen-go/OpenStars/Platform/Passport"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"github.com/OpenStars/thriftpool"	
	)


var (
	passportBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFunc( func( c thrift.TClient ) (interface{}) { return  (Passport.NewTPassportServiceClient(c)) }),
		thriftpool.DefaultClose)

	passportCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600, 
		thriftpool.GetThriftClientCreatorFuncCompactProtocol( func(c thrift.TClient) (interface{}) { return  (Passport.NewTPassportServiceClient(c)) }),
		thriftpool.DefaultClose)		
	
	)

 


func init(){
	fmt.Println("init thrift passportservice client ");
}

//GetPassportBinaryClient client by host:port
func GetPassportBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := passportBinaryMapPool.Get(aHost, aPort).Get()
	return client;
}

//GetPassportCompactClient get compact client by host:port
func GetPassportCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient{
	client, _ := passportCommpactMapPool.Get(aHost, aPort).Get()
	return client;
}