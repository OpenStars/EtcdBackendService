package String2Int64Service

type String2Int64ServiceIf interface {
	PutData(key string, value int64) error
	GetData(key string) (int64, error)
}
