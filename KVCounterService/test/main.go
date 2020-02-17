package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/KVCounterService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	kvcountersv := KVCounterService.NewKVCounterServiceModel("/aa/bb/",
		[]string{"127.0.0.1:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "10.60.68.102",
			Port:      "8883",
			ServiceID: "/aa/bbb",
		})
	v, err := kvcountersv.GetValue("pubkey2")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(v)
}
