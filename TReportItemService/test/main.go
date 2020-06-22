package main

import (
	"log"

	ReportItemService "github.com/OpenStars/EtcdBackendService/TReportItemService"
)

func main() {
	service := ReportItemService.NewReportItemService([]string{"10.60.68.1.20:2379"}, "/test/", "10.60.68.103", "8883")
	item, err := service.GetData(2)
	log.Println("item", item, "err", err)
}
