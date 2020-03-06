package main

import (
	uidserviceclient "TrustKeys/SocialNetworks/Account/UIDService/client"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/OpenStars/EtcdBackendService/Int2StringService"
	"github.com/OpenStars/EtcdBackendService/String2Int64Service"

	"github.com/OpenStars/EtcdBackendService/KVCounterService"
	"github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"

	tpostserviceclient "github.com/OpenStars/EtcdBackendService/TPostStorageService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func DumpData() {
	GenIDPost := "|GenIDPost|"
	mapPubkey2Uid := make(map[string]int64)
	kvcounter := KVCounterService.NewKVCounterServiceModel("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.60.68.103",
		Port:      "7974",
		ServiceID: "/test/",
	})
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.60.68.102",
		Port:      "8513",
		ServiceID: "/test/",
	})

	// 	[int2string]
	// sid = /openstars/trustkeys/cryptopayment/s2i64kv/
	// host = 10.60.68.103
	// port = 27183
	// [string2int]
	// sid =  /openstars/trustkeys/cryptopayment/s2i64kv
	// host = 10.60.68.103
	// port =  27173

	aint2string := Int2StringService.NewInt2StringService("/test/", []string{"10.60.1.20:2378"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.68.103",
		Port:      "27183",
		ServiceID: "/test/",
	})
	astring2int := String2Int64Service.NewString2Int64Service("/test/", []string{"10.60.1.20:2378"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.68.103",
		Port:      "27173",
		ServiceID: "/test/",
	})

	// uidservice := uidserviceclient.NewUIDServiceClient("10.60.68.103", "12010")
	maxidPost, _ := kvcounter.GetCurrentValue(GenIDPost)
	for i := int64(0); i <= maxidPost; i++ {
		postItem, err := tpostclient.GetData(i)
		if err != nil {
			continue
		}
		_, ok := mapPubkey2Uid[postItem.Pubkey]
		if !ok {
			if postItem.UID != 0 {
				mapPubkey2Uid[postItem.Pubkey] = postItem.UID
				log.Println("PostID", postItem.Idpost, "Pubkey", postItem.Pubkey, "UID", postItem.UID)
				if postItem.UID == 1205 {
					log.Println("UID = 1205")
					fmt.Scan()
				}
				err := aint2string.PutData(postItem.UID, postItem.Pubkey)
				if err != nil {
					log.Println("Int2String putdata err", err)
				}
				err = astring2int.PutData(postItem.Pubkey, postItem.UID)
				if err != nil {
					log.Println("String2Int putdata err", err)
				}
				// r, err := uidservice.SetMapUIDByPubkey(postItem.UID, postItem.Pubkey)
				// log.Println("Set map pubkey", postItem.Pubkey, "=> uid", postItem.UID, "result", r, "err", err)
			}

		}
	}

}

func GetInt2String() {
	aint2string := Int2StringService.NewInt2StringService("/test/", []string{"10.60.1.20:2378"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.68.103",
		Port:      "27183",
		ServiceID: "/test/",
	})
	// pubkey, err := aint2string.GetData(1205)
	// log.Println("pubkey", pubkey, "err", err)
	for i := int64(0); i < 4000; i++ {
		data, err := aint2string.GetData(i)
		if err != nil {
			continue
		}
		log.Println("uid =", i, "pub=", data)
	}
}

func PutToUIDService() {
	uidservice := uidserviceclient.NewUIDServiceClient("127.0.0.1", "12010")
	// pubkey, err := uidservice.GetPubkeyByUID(1205)
	// log.Println("pubkey", pubkey, "err", err)
	// uidservice.SetMapUIDByPubkey(1205, "034e095ae7adcc9d102a8c2d606c0c7150353fc7d3b60924935abccdc12432f78b")
	// pubkey, _ := uidservice.GetPubkeyByUID(5)
	// log.Println("pubkey", pubkey)
	// uid, _ := uidservice.GetUIDByPubkey(pubkey)
	// log.Println("uid", uid)
	// for i := 0; i < 1000; i++ {
	// 	uidmax, err := uidservice.GetMaxUID()
	// 	if err != nil {
	// 		log.Println("err")
	// 	} else {
	// 		log.Println("uidmax", uidmax)
	// 	}
	// }

	for i := int64(0); i < 4000; i++ {
		pubkey, err := uidservice.GetPubkeyByUID(i)
		if i%1000 == 0 {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}

		// log.Println("pubkey ", pubkey, "uid", i, "err", err)
		if err != nil {
			log.Println("i", i, "err", err)
			continue
		}

		log.Println(i, "pubkey=", pubkey)
	}

	err := make(chan bool)
	<-err

}

func TestPut() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"127.0.0.1"}, GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "8883",
		ServiceID: "/test/",
	})
	err := tpostclient.PutData(1, &TPostStorageService.TPostItem{
		Idpost:     1,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 1 err", err)
		return
	}
	err = tpostclient.PutData(2, &TPostStorageService.TPostItem{
		Idpost:     2,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 2 err", err)
		return
	}
	err = tpostclient.PutData(3, &TPostStorageService.TPostItem{
		Idpost:     3,
		Content:    "Xin chao moi nguoi",
		UID:        1,
		Timestamps: time.Now().Unix(),
	})
	if err != nil {
		log.Println("Put 3 err", err)
		return
	}
	log.Println("Put oke")
}

func TestGet() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"10.60.1.20:2379"}, GoEndpointBackendManager.EndPoint{
		Host:      "10.60.68.102",
		Port:      "8513",
		ServiceID: "/test/",
	})
	listdata, err := tpostclient.GetData(21191)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("pubkey ", listdata.Pubkey)
	log.Println("uid", listdata.UID)
	log.Println(listdata.Content)
}

func TestRemove() {
	tpostclient := tpostserviceclient.NewTPostStorageService("/test/", []string{"127.0.0.1"}, GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "8883",
		ServiceID: "/test/",
	})
	err := tpostclient.RemoveData(1)
	if err != nil {
		log.Println("Err:", err)
		return
	}
	data, err := tpostclient.GetData(1)
	if err != nil {
		log.Println("remove oke")
		return
	}
	log.Println(data)
}
func main() {
	// DumpData()
	PutToUIDService()
	// GetInt2String()

	// TestGet()
}
