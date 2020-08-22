package main

import (
	"log"

	tmediastorageclient "github.com/OpenStars/EtcdBackendService/TMediaStorageService"
)

// [0]:85072
// [1]:85030
// [2]:83775
// [3]:83569
// [4]:83538

func TestGet() {
	mediaclient := tmediastorageclient.NewTMediaStorageService2("/test/", []string{"10.60.1.20:2379"}, "10.110.68.103", "8973")
	data, err := mediaclient.GetListData([]int64{85072, 85030, 83775, 83569, 83538})
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("data", data)
}
func main() {
	TestGet()
}
