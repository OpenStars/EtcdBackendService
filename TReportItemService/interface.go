package ReportItemService

import "github.com/OpenStars/EtcdBackendService/TReportItemService/treportitemservice/thrift/gen-go/OpenStars/Platform/MarketPlace"

type ReportItemService interface {
	GetData(itemID int64) (*MarketPlace.TReportItem, error)
	PutData(itemID int64, data *MarketPlace.TReportItem) error
	RemoveData(itemID int64) error
	GetListDatas(listID []int64) ([]*MarketPlace.TReportItem, error)
}
