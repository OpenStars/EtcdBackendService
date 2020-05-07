package MarketPlaceService

import (
	"github.com/OpenStars/EtcdBackendService/MarketPlaceService/marketplaceitem/thrift/gen-go/OpenStars/Platform/MarketPlace"
)

type TMarketPlaceItemService interface {
	GetData(itemID int64) (*MarketPlace.TMarketPlaceItem, error)
	PutData(itemID int64, data *MarketPlace.TMarketPlaceItem) error
	RemoveData(itemID int64) error
	GetListDatas(listID []int64) ([]*MarketPlace.TMarketPlaceItem, error)
}
