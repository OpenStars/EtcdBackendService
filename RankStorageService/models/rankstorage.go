package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type Item struct {
	Key   []byte `json:"key,omitempty"`
	Value []byte `json:"value,omitempty"`
}

const (
	RANK_KEY_MAPPING_SCORE           = "RANK_KEY_MAPPING_SCORE_"
	RANK_KEY_MAPPING_ITEM            = "RANK_KEY_MAPPING_ITEM_"
	SCORE_MAPPING_ITEM               = "SCORE_MAPPING_ITEM"
	RANK_KEY_MAPPING_SCORE_WITH_ITEM = "RANK_KEY_MAPPING_SCORE_WITH_ITEM_"
	TEM_PADDING                      = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
)

type rankstoragemodels struct {
	rankKeyMappingScore      StringBigsetService.StringBigsetServiceIf
	rankKeyMappingItem       StringBigsetService.StringBigsetServiceIf
	scoreMappingItem         StringBigsetService.StringBigsetServiceIf
	rankMappingScoreWithItem StringBigsetService.StringBigsetServiceIf
	maxLenghtItemkey         int
}

func PadingZeros(num int64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func (m *rankstoragemodels) PaddingItemKey(itemkey string) string {
	if len(itemkey) < m.maxLenghtItemkey {
		return TEM_PADDING[:m.maxLenghtItemkey-len(itemkey)] + itemkey
	}
	return itemkey
}

// Lưu item ID với score vào 1 rankKey
func (m *rankstoragemodels) PutItemRank(rankKey string, data *Item, score int64) error {

	// put rankKey -> score -- lastest itemID
	err := m.rankKeyMappingScore.BsPutItem(generic.TStringKey(RANK_KEY_MAPPING_SCORE+rankKey), &generic.TItem{Key: []byte(PadingZeros(score)), Value: data.Key})
	if err != nil {
		return err
	}

	item, err := m.rankKeyMappingItem.BsGetItem(generic.TStringKey(RANK_KEY_MAPPING_ITEM+rankKey), generic.TItemKey(data.Key))
	if err != nil {
		return err
	}
	//  Nếu đã tồn tại item
	if item != nil {
		// remove old score of item ID
		oldscore := string(item.Value)
		bskey := generic.TStringKey(SCORE_MAPPING_ITEM + rankKey + "_" + oldscore)
		err = m.scoreMappingItem.BsRemoveItem(bskey, generic.TItemKey(data.Key))
		if err != nil {
			return err
		}
		// check if len list item of old score == 0 => remove score from rankKey mapping score
		total, _ := m.scoreMappingItem.GetTotalCount(bskey)

		if total == 0 {
			m.rankKeyMappingScore.BsRemoveItem(generic.TStringKey(RANK_KEY_MAPPING_SCORE+rankKey), generic.TItemKey(oldscore))
		}
		m.rankMappingScoreWithItem.BsRemoveItem(generic.TStringKey(RANK_KEY_MAPPING_SCORE_WITH_ITEM+rankKey), generic.TItemKey(oldscore+"_"+m.PaddingItemKey(string(data.Key))))
	}

	// put score mapping item
	newscore := PadingZeros(score)
	bskey := generic.TStringKey(SCORE_MAPPING_ITEM + rankKey + "_" + newscore)
	err = m.scoreMappingItem.BsPutItem(bskey, &generic.TItem{Key: data.Key, Value: data.Value})
	// rankKey -> item ID -- score
	err = m.rankKeyMappingItem.BsPutItem(generic.TStringKey(RANK_KEY_MAPPING_ITEM+rankKey), &generic.TItem{Key: data.Key, Value: []byte(PadingZeros(score))})
	if err != nil {
		return err
	}
	m.rankMappingScoreWithItem.BsPutItem(generic.TStringKey(RANK_KEY_MAPPING_SCORE_WITH_ITEM+rankKey), &generic.TItem{Key: []byte(newscore + "_" + m.PaddingItemKey(string(data.Key))), Value: item.Value})
	return nil
}

// Lấy top danh sách item theo score từ vị trí from với kích thước size , order = asc || desc tương ứng với lấy theo chiều tăng dần hoặc giảm dần của score
func (m *rankstoragemodels) GetTopItemRank(rankKey string, from int64, size int64, order string) ([]*Item, error) {
	var lsItem []*generic.TItem
	var err error
	if order == "asc" {
		lsItem, err = m.rankMappingScoreWithItem.BsGetSlice(generic.TStringKey(RANK_KEY_MAPPING_SCORE_WITH_ITEM+rankKey), int32(from), int32(size))
	} else {
		lsItem, err = m.rankMappingScoreWithItem.BsGetSliceR(generic.TStringKey(RANK_KEY_MAPPING_SCORE_WITH_ITEM+rankKey), int32(from), int32(size))
	}
	if err != nil {
		return nil, err
	}
	if lsItem == nil || len(lsItem) == 0 {
		return nil, nil
	}
	databytes, _ := json.Marshal(lsItem)
	var rs []*Item
	err = json.Unmarshal(databytes, &rs)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(rs); i++ {
		slpitkey := strings.Split(string(rs[i].Key), "_")
		if len(slpitkey) != 2 {
			continue
		}
		rs[i].Key = []byte(slpitkey[1])
	}
	return rs, nil
}
func (m *rankstoragemodels) GetTopItemRankFromScore(rankKey string, score int64, size int64, direction string) ([]*Item, error) {

}
func (m *rankstoragemodels) GetCurrentScoreOfItem(rankKey string, itemKey string) (int64, error) {

}

func (m *rankstoragemodels) RemoveItem(rankKey string, itemKey string) error {
	item, err := m.rankKeyMappingItem.BsGetItem(generic.TStringKey(RANK_KEY_MAPPING_ITEM+rankKey), generic.TItemKey(itemKey))
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("Can't found itemkey " + itemKey)
	}
	err = m.rankKeyMappingItem.BsRemoveItem(generic.TStringKey(RANK_KEY_MAPPING_ITEM+rankKey), generic.TItemKey(itemKey))
	err = m.scoreMappingItem.BsRemoveItem(generic.TStringKey(SCORE_MAPPING_ITEM+rankKey+"_"+string(item.Value)), generic.TItemKey(itemKey))
	total, _ := m.scoreMappingItem.GetTotalCount(generic.TStringKey(SCORE_MAPPING_ITEM + rankKey + "_" + itemKey))
	if total == 0 {
		err = m.rankKeyMappingScore.BsRemoveItem(generic.TStringKey(RANK_KEY_MAPPING_SCORE+rankKey), generic.TItemKey(item.Value))
	}
	if err != nil {
		return err
	}
	return nil
}
