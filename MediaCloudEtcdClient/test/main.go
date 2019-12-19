package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendSerivce/MediaCloudEtcdClient"
)

func Test() {
	mediacloudclient := MediaCloudEtcdClient.NewPubProfileClient("10.60.68.100", "9210")
	r, err := mediacloudclient.GetMediaInfo("9b5323711834f16aa825", "c486c2a5c1fb1f0f1071c29593f2f260", "cbba01983addd3838acc")
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("status", *r.Value.Status)
	for i := 0; i < len(r.Value.PosterUrls); i++ {
		for k, v := range r.Value.PosterUrls[i] {
			log.Println("k=", k)
			log.Println("v=", *v.FileUrl)
		}
	}
}
func main() {
	Test()
}
