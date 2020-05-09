package main

import (
	"fmt"

	"github.com/OpenStars/EtcdBackendService/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("127.0.0.1", "1805")
	profileData, _ := pubclient.GetProfileByPubkey("024ede744912fe25c5e684f259ddece762d7c1bad0c15ff2d6c536528e194db0fa")
	fmt.Println(profileData)
	// profileData.Pubkey = "030ca42cf118bd01574f52bf291e7260201ea73de51855af4d3487d9c9945aaf73"
	// profileData.LinkFB = "ahihi"
	// r, err := pubclient.UpdateProfileByPubkey("030ca42cf118bd01574f52bf291e7260201ea73de51855af4d3487d9c9945aaf73", profileData)
	// if err != nil {
	// 	log.Println("err", err)
	// 	return
	// }
	// log.Println(r)
}
func main() {
	Test()
}
