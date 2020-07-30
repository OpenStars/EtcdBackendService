package main

import (
	"log"

	TNotifyStorageService "github.com/OpenStars/EtcdBackendService/TNotifyStorageService2"
	notifistoragedata "github.com/OpenStars/EtcdBackendService/TNotifyStorageService2/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
)

func TestPut() {
	notificlient := TNotifyStorageService.NewTNotifyStorageService2(nil, "/test/", "127.0.0.1", "8883")
	data := &notifistoragedata.TNotifyItem{
		ID:         1,
		ActionType: "test",
	}
	ok, err := notificlient.PutData(data.ID, data)
	log.Println("oke", ok, "err", err)

}

func TestGet() {
	notificlient := TNotifyStorageService.NewTNotifyStorageService2(nil, "/test/", "127.0.0.1", "8883")
	data, err := notificlient.GetData(1)
	log.Println("data", data, "err", err)
}

func TestRemove() {

}
func main() {
	// TestPut()
	TestGet()
}
