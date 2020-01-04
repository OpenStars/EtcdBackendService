package TMediaStorageService

import (
	"github.com/OpenStars/backendclients/go/tmediastorageservice/thrift/gen-go/OpenStars/Common/TMediaStorageService"
)

type TMediaStorageServiceIf interface {
	GetData(idmedia int64) (*TMediaStorageService.TMediaItem, error)
	PutData(idmedia int64, data *TMediaStorageService.TMediaItem) error
	RemoveData(idmedia int64) error

	GetListData(listkey []int64) (r []*TMediaStorageService.TMediaItem, err error)
}
