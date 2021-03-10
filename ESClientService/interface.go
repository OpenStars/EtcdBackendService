package ESClientService

import "github.com/olivere/elastic"

type ESClientServiceIf interface {
	PutDataToES(id string, dataJson string) (err error)
	DeleteIndexES()
	DeleteDataES(id string)
	UpdateDataES(id string, mapUpdate map[string]interface{})
	PutDataToES2(id string, data interface{}) error
	PutDataToES3(data interface{}) error
	SearchESByQuery(mapSearch map[string]interface{}, sort map[string]bool) ([]*elastic.SearchHit, error)
	GetClientES() *elastic.Client
}
