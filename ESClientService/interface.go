package ESClientService

type ESClientServiceIf interface {
	PutDataToES(id string, dataJson string) (err error)
	DeleteIndexES()
	DeleteDataES(id string)
	UpdateDataES(id string, mapUpdate map[string]interface{})
	PutDataToES2(id string, data interface{}) error
}
