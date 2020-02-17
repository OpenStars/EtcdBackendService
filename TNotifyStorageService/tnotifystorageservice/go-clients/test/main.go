package main
import (
	"../../thrift/gen-go/OpenStars/Common/TNotifyStorageService" //Todo: Fix this
	"../transports"  //Todo: Fix this
	"fmt"
	"context"
)

func main(){
	aClient := transports.GetTNotifyStorageServiceCompactClient("127.0.0.1", "8883") //Todo: Check port and protocol
	defer aClient.BackToPool();
	fmt.Println("Client: ", aClient);

	res, _ := aClient.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetData(context.Background(), (TNotifyStorageService.TKey)(10) )
	fmt.Println("get result: ",res)

	aClient.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).PutData(context.Background(),
			 (TNotifyStorageService.TKey)(10), &TNotifyStorageService.TNotifyItem{} );//Todo: fill structure here

	res, _ = aClient.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetData(context.Background(), (TNotifyStorageService.TKey)(10) )
	fmt.Println("get after put :", res)

}