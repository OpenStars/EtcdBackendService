package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/OpenStars/EtcdBackendService/ESClientService"
	"github.com/elastic/go-elasticsearch"
)

type MarketPlaceItem struct {
	PostID  int64  `json:"postID"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Price   int64  `json:"price"`
}

func TestSearch() {
	//	esclient := ESClientService.NewESClient("http://localhost:9200/", "marketplace", "marketplaceitem")
	// esclient.SearchESByQuery(map[string]interface{}{
	// 	"query" :
	// })
	// client, err := elastic.NewClient(elastic.SetURL("http://10.110.1.100:9206/"),
	// 	elastic.SetSniff(false),
	// 	elastic.SetHealthcheck(false))
	// // termQuery := elastic.NewTermQuery("title", "x2 carbon")
	// // marchQuery := elastic.NewMatchQuery("title", "x2 carbon")
	// // marchQuery := elastic.NewMatchPhraseQuery("title", "car")
	// // suggest := elastic.NewSuggestField("title", "car")

	// marchQuery := elastic.("x2", "query")

	// rs, err := client.Search().Index("marketplace").Query(marchQuery).Do(context.Background())

	// for _, item := range rs.Hits.Hits {
	// 	data, err := item.Source.MarshalJSON()
	// 	if err != nil {
	// 		log.Println("err", err)
	// 	}
	// 	log.Println("data", string(data))
	// }
	// log.Println("rs", rs, "err", err)
}
func TestPut() {
	esclient := ESClientService.NewESClient("http://localhost:9200/", "marketplace", "marketplaceitem")

	item1 := &MarketPlaceItem{
		PostID:  4,
		Content: "X2 Carbon i7 ram 8gb ssd 256",
		Title:   "x2 carbon",
		Price:   2000000,
	}

	// item2 := &MarketPlaceItem{
	// 	PostID:  2,
	// 	Content: "Xiaomi mi8 lite ram 4gb snapdragon 660",
	// 	Title:   "Xiaomi mi8 lite",
	// 	Price:   7000000,
	// }

	// item3 := &MarketPlaceItem{
	// 	PostID:  1,
	// 	Content: "Honda civic 2020 4x4 2.0l AT",
	// 	Title:   "Honda civic 2020",
	// 	Price:   800000000,
	// }

	err := esclient.PutDataToES2("4", item1)
	if err != nil {
		log.Println("item1 err", err)
	}
	// err = esclient.PutDataToES2("2", item2)
	// if err != nil {
	// 	log.Println("item2 err", err)
	// }

	// err = esclient.PutDataToES2("3", item3)
	// if err != nil {
	// 	log.Println("item3 err", err)
	// }
}

func OfficalSearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://10.110.1.100:9206/",
		},
	}
	esclient, err := elasticsearch.NewClient(cfg)
	log.Println("esclient", esclient, "err", err)
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": "thinkpad",
				// "fields": []string{"_all"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esclient.Search(
		esclient.Search.WithContext(context.Background()),
		esclient.Search.WithIndex("marketplace"),
		esclient.Search.WithBody(&buf),
		esclient.Search.WithTrackTotalHits(true),
		esclient.Search.WithPretty(),
	)
	if err != nil {
		log.Println("err", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalln("status", res.Status, "err", err)
		}
	}
	datastring, err := ioutil.ReadAll(res.Body)
	log.Println("data string", string(datastring), "err", err)

}

func main() {
	// TestSearch()
	// TestPut()
	OfficalSearch()
}
