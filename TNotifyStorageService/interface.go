package TPostStorageService

import "github.com/OpenStars/backendclients/go/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"

type TNotifyStorageServiceIf interface {
	GetData(idpost int64) (*TNotifyStorageService.TNotifyItem, error)
	PutData(idpost int64, data *TNotifyStorageService.TNotifyItem) error
	RemoveData(idpost int64) error
	GetListDatas(listkey []int64) ([]*TNotifyStorageService.TNotifyItem, error)
}
