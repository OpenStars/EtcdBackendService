package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/MarketPlaceService"
)

func main() {
	marketplaceservice := MarketPlaceService.NewTMarketPlaceItemService(nil, "/test/", "127.0.0.1", "8883")
	// err := marketplaceservice.PutData(0, &MarketPlace.TMarketPlaceItem{
	// 	ID:    0,
	// 	Price: 50000,
	// 	Title: "Xiaomi redmi note 8",
	// })
	// err = marketplaceservice.PutData(1, &MarketPlace.TMarketPlaceItem{
	// 	ID:    0,
	// 	Price: 50000,
	// 	Title: "Xiaomi redmi 8 lite",
	// })
	err := marketplaceservice.RemoveData(1)
	if err != nil {
		log.Println("err", err)
	} else {
		// log.Println("item", item)
	}
	lsItem, err := marketplaceservice.GetListDatas([]int64{0, 1})
	if err != nil {
		log.Println("err", err)
	} else {
		log.Println("lsitems", lsItem)
	}
}
