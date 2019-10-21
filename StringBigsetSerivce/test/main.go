package main

import (
	"log"
	"strconv"

	"github.com/OpenStars/EtcdBackendSerivce/StringBigsetSerivce"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

func TestSV() {
	svClient := StringBigsetSerivce.NewStringBigsetServiceModel("/aa/bb", []string{"127.0.0.1:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "127.0.0.1",
			Port:      "18407",
			ServiceID: "/aa/bb",
		})

	for i := 0; i < 30; i++ {
		item := &generic.TItem{
			Key:   []byte(strconv.FormatInt(int64(i), 10)),
			Value: []byte("test"),
		}
		svClient.BsPutItem(generic.TStringKey("test"), item)

	}

	startKey := generic.TItemKey(strconv.FormatInt(1, 10))
	endKey := generic.TItemKey(strconv.FormatInt(5, 10))
	item, _ := svClient.BsRangeQuery(generic.TStringKey("test"), startKey, endKey)
	for i := 0; i < len(item); i++ {
		log.Println(string(item[i].Key))
	}

}
func main() {
	TestSV()
}
