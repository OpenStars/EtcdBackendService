package TNotifyStorageService

import (
	"github.com/OpenStars/EtcdBackendService/TNotifyStorageService/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
)

type TNotifyStorageServiceIf interface {
	GetData(idnotify int64) (*TNotifyStorageService.TNotifyItem, error)
	PutData(idnotify int64, data *TNotifyStorageService.TNotifyItem) error
	RemoveData(idnotify int64) error
	GetListDatas(listkey []int64) ([]*TNotifyStorageService.TNotifyItem, error)
}
