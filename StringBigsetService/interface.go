package StringBigsetService

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type StringBigsetServiceIf interface {
	TotalStringKeyCount() (r int64, err error)
	GetListKey(fromIndex int64, count int32) ([]string, error)
	BsPutItem(bskey generic.TStringKey, item *generic.TItem) error
	BsGetItem(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error)
	GetTotalCount(bskey generic.TStringKey) (int64, error)
	BsMultiPut(bskey generic.TStringKey, lsItems []*generic.TItem) error
	BsGetSlice(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
	BsRemoveItem(bskey generic.TStringKey, itemkey generic.TItemKey) error
	GetBigSetInfoByName(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	CreateStringBigSet(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	BsRangeQuery(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error)
	BsRangeQueryByPage(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error)
	BsGetSliceFromItem(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	BsGetSliceFromItemR(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	RemoveAll(bskey generic.TStringKey) error
	BsGetSliceR(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
}
