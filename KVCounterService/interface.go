package KVCounterService

type KVCounterServiceIf interface {
	GetValue(genname string) (int64, error)
	CreateGenerator(genname string) (int32, error)
}
