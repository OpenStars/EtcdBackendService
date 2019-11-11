package StringMapKV

type StringMapKVIf interface {
	GetData(key string) (string, error)
	PutData(key string, value string) error
	DeleteKey(key string) error
}
