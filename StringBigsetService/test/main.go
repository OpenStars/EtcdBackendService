package main

import (
	"TrustKeys/SocialNetworks/Centerhub/model/share"
	"TrustKeys/SocialNetworks/Centerhub/util"
	"log"

	"github.com/OpenStars/EtcdBackendSerivce/StringBigsetSerivce"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

func TestSV() {
	svClient := StringBigsetSerivce.NewStringBigsetServiceModel("/aa/bb", []string{"127.0.0.1:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "10.60.68.102",
			Port:      "20517",
			ServiceID: "/aa/bb",
		})

	lsItem, err := svClient.BsGetSlice(generic.TStringKey(share.UIDPostPrefix+util.PadingZerors(911)), 2, 5)
	if err != nil {
		log.Println("err", err)
		return
	}
	// log.Println("total", total)
	for _, item := range lsItem {
		log.Println(string(item.Key))
	}

	// for i := 0; i < 30; i++ {
	// 	item := &generic.TItem{
	// 		Key:   []byte(strconv.FormatInt(int64(i+100), 10)),
	// 		Value: []byte("test"),
	// 	}
	// 	err := svClient.BsPutItem(generic.TStringKey("test"), item)
	// 	if err != nil {
	// 		log.Println("bsputitem err", err)
	// 	}
	// 	// reader := bufio.NewReader(os.Stdin)
	// 	// fmt.Print("Enter text: ")
	// 	// text, _ := reader.ReadString('\n')
	// 	// fmt.Println(text)
	// }

	// lsItem, err := svClient.BsGetSliceFromItemR(generic.TStringKey("test"), generic.TItemKey(strconv.FormatInt(120, 10)), 10)
	// if err != nil {
	// 	log.Println("svClient err", err)
	// }
	// for _, item := range lsItem {
	// 	log.Println(string(item.Key))
	// }

	// startKey := generic.TItemKey(strconv.FormatInt(1, 10))
	// endKey := generic.TItemKey(strconv.FormatInt(5, 10))
	// item, _ := svClient.BsRangeQuery(generic.TStringKey("test"), startKey, endKey)
	// for i := 0; i < len(item); i++ {
	// 	log.Println(string(item[i].Key))
	// }

}
func main() {
	TestSV()
}
