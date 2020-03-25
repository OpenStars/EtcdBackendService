package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("127.0.0.1", "1805")
	profileData, _ := pubclient.GetProfileByPubkey("0343038f99f26c910b4e0b0775d25d87ba1d57e2ef7a7fb34f777c356964bf88a2")
	profileData.DisplayName = "Nguyễn Thị Kim Liên"
	r, err := pubclient.UpdateProfileByPubkey("0343038f99f26c910b4e0b0775d25d87ba1d57e2ef7a7fb34f777c356964bf88a2", profileData)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println(r)
}
func main() {
	Test()
}
