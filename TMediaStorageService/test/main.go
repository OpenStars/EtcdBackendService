package main

import (
	"log"

	tmediastorageclient "github.com/OpenStars/EtcdBackendSerivce/TMediaStorageService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/tmediastorageservice/thrift/gen-go/OpenStars/Common/TMediaStorageService"
)

func TestPut() {
	mediaclient := tmediastorageclient.NewTMediaStorageService("/test/", []string{"127.0.0.1:2379"}, GoEndpointBackendManager.EndPoint{
		Host: "10.60.68.102",
		Port: "8973",
		// Host:      "127.0.0.1",
		// Port:      "8883",
		ServiceID: "/test/",
	})
	err := mediaclient.PutData(1, &TMediaStorageService.TMediaItem{
		Idmedia: 1,
		URL:     "http://testurl.html",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("oke")
}
func main() {
	TestPut()
}
