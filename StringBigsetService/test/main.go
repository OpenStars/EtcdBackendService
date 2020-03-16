package main

import (
	"TrustKeys/SocialNetworks/Account/UIDService/client"
	"TrustKeys/SocialNetworks/Centerhub/model/share"
	"TrustKeys/SocialNetworks/Centerhub/util"
	"log"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func TestSV() {
	svClient := StringBigsetService.NewStringBigsetServiceModel("/trustkeys/socialnetwork/newsfeed/stringbs", []string{"10.60.1.20:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "10.60.68.102",
			Port:      "20517",
			ServiceID: "/aa/bb",
		})

	uidservice := client.NewUIDServiceClient("10.60.68.103", "12010")
	uid, err := uidservice.GetUIDByPubkey("02898dd812414d661b7b9c0dee015ef6e3a92c943cfb64e838f14ff093ad9dd93f")
	if err != nil {
		log.Println("err ", err)
		return
	}
	lsItem, err := svClient.BsGetSlice(generic.TStringKey(share.NewsFeedPrefix+util.PadingZerors(uid)), 0, 1000)
	if err != nil {
		log.Println("err", err)
		return
	}
	// log.Println("total", total)
	for _, item := range lsItem {
		log.Println("postID", string(item.Key))
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
