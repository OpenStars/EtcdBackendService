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
	for f, mapquality := range r.Value.MapFormatQualityInfo {
		log.Println("Format", f)
		log.Println("Value", mapquality)
		for q, filestatusurl := range mapquality {
			log.Println("quality", q)
			log.Println("filestatusurl", filestatusurl)
		}
	}
}
func main() {
	Test()
}
