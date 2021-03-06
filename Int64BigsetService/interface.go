package Int64BigsetService

import "github.com/OpenStars/EtcdBackendService/Int64BigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"

type Int64BigsetServiceIf interface {
	PutItem(bskey generic.TKey, item *generic.TItem) error
	GetItem(bskey generic.TKey, itemkey generic.TItemKey) (*generic.TItem, error)
	GetTotalCount(bskey generic.TKey) (int64, error)
	GetSlice(bskey generic.TKey, fromPos int32, count int32) ([]*generic.TItem, error)
	MultiPut(bskey generic.TKey, lsItems []*generic.TItem) error
	RemoveItem(bskey generic.TKey, itemkey generic.TItemKey) error
	RangeQuery(bskey generic.TKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error)
}
