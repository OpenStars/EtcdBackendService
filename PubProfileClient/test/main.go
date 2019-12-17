package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendSerivce/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("10.60.1.20", "1805")
	r, err := pubclient.GetProfileByUID(5)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println(r)
}
func main() {
	Test()
}
