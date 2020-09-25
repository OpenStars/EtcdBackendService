package vtpcmtorderservice

import (
	"github.com/OpenStars/EtcdBackendService/vtpcmtorderservice/thrift/gen-go/OpenStars/VTPComment"
)

type VTPCommentOrderService interface {
	GetData(key string) (*VTPComment.TVTPComment, error)
	GetMultiData(keys []VTPComment.TKey) (map[VTPComment.TKey]*VTPComment.TVTPComment, error)
	PutData(lognumber string, data *VTPComment.TVTPComment) error
}
