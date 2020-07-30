package main

import (
	"encoding/json"
	"log"

	"github.com/OpenStars/backendclients/go/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
)

func TestPut() {
	// notificlient := notifistorageservie.NewTNotifyStorageService("/test/", []string{"127.0.0.1:2379"}, GoEndpointBackendManager.EndPoint{
	// 	Host:      "10.60.68.102",
	// 	Port:      " 8533",
	// 	ServiceID: "/test/",
	// })
	data := &TNotifyStorageService.TNotifyItem{
		Key:       11,
		ActionId:  5,
		ObjectId:  6,
		SubjectId: 7,
		MapData:   make(map[string]string),
	}
	data.MapData["xyz"] = "abc"

	var newdata TNotifyStorageService.TNotifyItem

	databytes, _ := json.Marshal(data)

	json.Unmarshal(databytes, &newdata)

	newdatabytes, _ := json.Marshal(newdata)

	log.Println(string(newdatabytes))
	// notificlient.PutData(11, data)
	// item, _ := notificlient.GetData(11)
	// log.Println(item)
}

func TestRemove() {

}
func main() {
	TestPut()
}
