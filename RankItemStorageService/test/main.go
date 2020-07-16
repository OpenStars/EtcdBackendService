package main

import (
	"TrustKeys/RankItemStorageService/models"
	"TrustKeys/TKRealtimeClient/common/util"
	"math/rand"
	"time"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
)

func main() {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", "127.0.0.1", "19417")
	md := models.NewRankItemScore(stringbs, 1000, 19)
	rand.Seed(time.Now().Unix())
	for i := int64(0); i < 1000; i++ {
		rankscore := rand.Int63n(10)
		md.PutItemScore("TEST", util.PaddingZeros(i), "oke", rankscore)
	}

}
