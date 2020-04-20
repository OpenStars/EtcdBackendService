package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func TestSV() {
	svClient := StringBigsetService.NewStringBigsetServiceModel("/test/", []string{"10.60.1.20:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "10.60.68.103",
			Port:      "20507",
			ServiceID: "/aa/bb",
		})

	total, err := svClient.TotalStringKeyCount()
	// lsItems, err := svClient.TotalStringKeyCount(generic.TStringKey("02ea252935dfc60ed6c882897ee0a52b6ac30fa5fa7570344317e6a5e6ef52a87f"), 0, 1)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("totala", total)
	// total, err := svClient.TotalStringKeyCount()
	// lsKey, err := svClient.GetListKey(0, int32(total))
	// if err != nil {
	// 	log.Fatalln("GetListKey err", err)
	// }
	// for _, key := range lsKey {
	// 	log.Println("key", key)
	// }
	// uidservice := client.NewUIDServiceClient("10.60.68.103", "12010")
	// uid, err := uidservice.GetUIDByPubkey("02898dd812414d661b7b9c0dee015ef6e3a92c943cfb64e838f14ff093ad9dd93f")
	// if err != nil {
	// 	log.Println("err ", err)
	// 	return
	// }

	// lsItem, err := svClient.BsGetSlice(generic.TStringKey(share.NewsFeedPrefix+util.PadingZerors(uid)), 0, 780)
	// if err != nil {
	// 	log.Println("err", err)
	// 	return
	// }

	// total, err := svClient.GetTotalCount(generic.TStringKey(share.NewsFeedPrefix + util.PadingZerors(uid)))
	// log.Println("bskey", string(generic.TStringKey(share.NewsFeedPrefix+util.PadingZerors(uid))))
	// if err != nil {
	// 	log.Println("Total err", err)
	// 	return
	// }
	// lsItem, err := svClient.BsGetSlice(generic.TStringKey(share.NewsFeedPrefix+util.PadingZerors(uid)), 0, int32(total))
	// for i := 0; i < len(lsItem); i++ {
	// 	log.Println("key", string(lsItem[i].Key))
	// }
	// log.Println("total", total)
	// log.Println("total", total)
	// for i := int32(0); i < int32(total); i++ {
	// 	lsItem, err := svClient.BsGetSlice(generic.TStringKey(share.NewsFeedPrefix+util.PadingZerors(uid)), i, 1)
	// 	if err != nil {
	// 		log.Println("GetItemkey", i, "err", err)
	// 	}
	// 	for _, item := range lsItem {
	// 		log.Println("postID", string(item.Key))
	// 	}

	// }

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
