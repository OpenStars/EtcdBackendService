package TNotifyStorageService

import (
	"github.com/OpenStars/EtcdBackendService/TNotifyStorageService2/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
)

type TNotifyStorageServiceIf interface {
	GetData(idnotify int64) (*TNotifyStorageService.TNotifyItem, error)
	PutData(idnotify int64, data *TNotifyStorageService.TNotifyItem) (bool, error)
	RemoveData(idnotify int64) (bool, error)
	GetListDatas(listkey []int64) ([]*TNotifyStorageService.TNotifyItem, error)
}
