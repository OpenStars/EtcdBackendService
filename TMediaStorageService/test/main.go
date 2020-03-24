package main

import (
	"log"

	tmediastorageclient "github.com/OpenStars/EtcdBackendService/TMediaStorageService"
)

func TestGet() {
	mediaclient := tmediastorageclient.NewTMediaStorageService2("/test/", []string{"10.60.1.20:2379"}, "10.60.68.102", "8973")
	data, err := mediaclient.GetData(22294)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("data", data)
}
func main() {
	TestGet()
}
