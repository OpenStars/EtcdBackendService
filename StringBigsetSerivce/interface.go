package StringBigsetSerivce

import (
	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type StringBigsetServiceIf interface {
	BsPutItem(bskey generic.TStringKey, item *generic.TItem) error
	BsGetItem(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error)
	GetTotalCount(bskey generic.TStringKey) (int64, error)
	BsMultiPut(bskey generic.TStringKey, lsItems []*generic.TItem) error
	BsGetSlice(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
	BsRemoveItem(bskey generic.TStringKey, itemkey generic.TItemKey) error
	GetBigSetInfoByName(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	CreateStringBigSet(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	BsRangeQuery(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error)
	BsGetSliceFromItem(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
}
