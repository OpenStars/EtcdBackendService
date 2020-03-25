package main

import (
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("127.0.0.1", "1805")
	profileData, _ := pubclient.GetProfileByPubkey("0344dd9748a1c687e91a1290bb5964d2b6b9803e2c219e9906915790add411ebf4")
	fmt.Println(profileData)
	profileData.Pubkey = "0344dd9748a1c687e91a1290bb5964d2b6b9803e2c219e9906915790add411ebf4"
	profileData.DisplayName = "Nguyễn Thị Kim Liên"
	r, err := pubclient.UpdateProfileByPubkey("0344dd9748a1c687e91a1290bb5964d2b6b9803e2c219e9906915790add411ebf4", profileData)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println(r)
}
func main() {
	Test()
}
