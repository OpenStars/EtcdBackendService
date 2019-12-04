package main

import (
	"log"
	"time"

	tpostserviceclient "github.com/OpenStars/EtcdBackendSerivce/TPostStorageService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
)

func TestPut() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"127.0.0.1"}, GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "8883",
		ServiceID: "/test/",
	})
	err := tpostclient.PutData(1, &TPostStorageService.TPostItem{
		Idpost:     1,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 1 err", err)
		return
	}
	err = tpostclient.PutData(2, &TPostStorageService.TPostItem{
		Idpost:     2,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 2 err", err)
		return
	}
	err = tpostclient.PutData(3, &TPostStorageService.TPostItem{
		Idpost:     3,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 3 err", err)
		return
	}
	log.Println("Put oke")
}

func TestGet() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"127.0.0.1"}, GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "8883",
		ServiceID: "/test/",
	})
	listdata, err := tpostclient.GetData(1)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println(listdata.Content)
}

func TestRemove() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"127.0.0.1"}, GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "8883",
		ServiceID: "/test/",
	})
	err := tpostclient.RemoveData(1)
	if err != nil {
		log.Println("Err:", err)
		return
	}
	data, err := tpostclient.GetData(1)
	if err != nil {
		log.Println("remove oke")
		return
	}
	log.Println(data)
}
func main() {
	TestRemove()
}
