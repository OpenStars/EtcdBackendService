package main

import "github.com/OpenStars/EtcdBackendService/ESClientService"

type MarketPlaceItem struct {
	PostID  int64  `json:"postID"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Price   int64  `json:"price"`
}

func main() {
	esclient := ESClientService.NewESClient("http://localhost:9200/", "marketplace", "marketplaceitem")
	err := esclient.PutDataToES("1")
}
