package TPostStorageService

import "github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"

type TPostStorageServiceIf interface {
	GetData(idpost int64) (*TPostStorageService.TPostItem, error)
	PutData(idpost int64, data *TPostStorageService.TPostItem) error
	RemoveData(idpost int64) error
	GetListDatas(listkey []int64) ([]*TPostStorageService.TPostItem, error)
	GetData2(idpost int64) (*TPostStorageService.TPostItem, error)
}
