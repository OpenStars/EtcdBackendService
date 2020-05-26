package main

import (
	"fmt"

	"github.com/OpenStars/EtcdBackendService/PubProfileClient"
)

func Test() {
	pubclient := PubProfileClient.NewPubProfileClient("10.60.68.102", "1805")
	profileData, e := pubclient.GetProfileByPubkey("minh")
	fmt.Println(profileData, e)
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
