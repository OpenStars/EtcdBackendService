package models

type RankStorage interface {
	PutItemRank(rankKey string, item Item, score int64) error
	GetTopItemRank(rankKey string, from int64, size int64, order string) ([]*Item, error)
	GetTopItemRankFromScore(rankKey string, score int64, size int64, direction string) ([]*Item, error)
	GetCurrentScoreOfItem(rankKey string, itemKey string) (int64, error)
}
