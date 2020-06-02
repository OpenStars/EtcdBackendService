package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/MarketPlaceService"
	"github.com/OpenStars/EtcdBackendService/MarketPlaceService/marketplaceitem/thrift/gen-go/OpenStars/Platform/MarketPlace"
)

func main() {
	marketplaceservice := MarketPlaceService.NewTMarketPlaceItemService(nil, "/test/", "10.60.68.103", "6003")
	err := marketplaceservice.PutData(21, &MarketPlace.TMarketPlaceItem{
		ID:    21,
		Price: 50000,
		Title: "Xiaomi redmi note 8",
		MapExtend: map[string]string{
			"KeyA": "A",
		},
	})
	log.Println("err", err)
	// err = marketplaceservice.PutData(1, &MarketPlace.TMarketPlaceItem{
	// 	ID:    0,
	// 	Price: 50000,
	// 	Title: "Xiaomi redmi 8 lite",
	// })
	// err := marketplaceservice.RemoveData(1)
	// if err != nil {
	// 	log.Println("err", err)
	// } else {
	// 	// log.Println("item", item)
	// }
	// lsItem, err := marketplaceservice.GetListDatas([]int64{0, 1})
	// if err != nil {
	// 	log.Println("err", err)
	// } else {
	// 	log.Println("lsitems", lsItem)
	// }
	log.Println(marketplaceservice.GetData(21))
}
