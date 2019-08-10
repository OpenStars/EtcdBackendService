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
}
