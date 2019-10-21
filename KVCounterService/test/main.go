package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendSerivce/KVCounterService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	kvcountersv := KVCounterService.NewStringBigsetServiceModel("/aa/bb/",
		[]string{"127.0.0.1:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "127.0.0.1",
			Port:      "2483",
			ServiceID: "/aa/bbb",
		})
	v, err := kvcountersv.GetValue("pubkey2")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(v)
}
