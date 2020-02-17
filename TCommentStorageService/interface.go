package TCommentStorageService

import "github.com/OpenStars/EtcdBackendService/TCommentStorageService/tcommentstorageservice/thrift/gen-go/OpenStars/Common/TCommentStorageService"

type TCommentStorageServiceIf interface {
	GetData(idcomment int64) (*TCommentStorageService.TCommentItem, error)
	PutData(idcomment int64, data *TCommentStorageService.TCommentItem) error
	RemoveData(idcomment int64) error
	GetListDatas(listkey []int64) ([]*TCommentStorageService.TCommentItem, error)
}
