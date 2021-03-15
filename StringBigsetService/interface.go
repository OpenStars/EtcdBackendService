package StringBigsetService

import (
	"github.com/Sonek-HoangBui/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
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
	BsMultiPutIndex(bskey []generic.TStringKey, lsItems []*generic.TItem) error

	//============================================================= Version 2 =========================================================================

	TotalStringKeyCount2() (r int64, err error)
	GetListKey2(fromIndex int64, count int32) ([]string, error)
	BsPutItem2(bskey generic.TStringKey, item *generic.TItem) (bool, error)
	BsGetItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error)
	GetTotalCount2(bskey generic.TStringKey) (int64, error)
	BsMultiPut2(bskey generic.TStringKey, lsItems []*generic.TItem) (bool, error)
	BsGetSlice2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
	BsRemoveItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (bool, error)
	GetBigSetInfoByName2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	CreateStringBigSet2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	BsRangeQuery2(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error)
	BsRangeQueryByPage2(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error)
	BsGetSliceFromItem2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	BsGetSliceFromItemR2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	RemoveAll2(bskey generic.TStringKey) (bool, error)
	BsGetSliceR2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)

	BsPutItemWithoutPutBackup(bskey generic.TStringKey, item *generic.TItem) error
}

type Client interface {
	TotalStringKeyCount2() (r int64, err error)
	GetListKey2(fromIndex int64, count int32) ([]string, error)
	BsPutItem2(bskey generic.TStringKey, item *generic.TItem) (bool, error)
	BsGetItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error)
	GetTotalCount2(bskey generic.TStringKey) (int64, error)
	BsMultiPut2(bskey generic.TStringKey, lsItems []*generic.TItem) (bool, error)
	BsGetSlice2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
	BsRemoveItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (bool, error)
	GetBigSetInfoByName2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	CreateStringBigSet2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error)
	BsRangeQuery2(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error)
	BsRangeQueryByPage2(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error)
	BsGetSliceFromItem2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	BsGetSliceFromItemR2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error)
	RemoveAll2(bskey generic.TStringKey) (bool, error)
	BsGetSliceR2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error)
}
