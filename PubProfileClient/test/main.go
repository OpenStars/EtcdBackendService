package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("10.60.68.102", "1805")
	r, err := pubclient.GetProfileByPubkey("039b81844d2caaf2b32f286b70c00db4c6309a10ccb8c4804e0f7d5df6f68e1a05")
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println(r)
}
func main() {
	Test()
}
