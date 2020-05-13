package Tile38ClientService

type Tile38ManagerServiceIf interface {
	DeleteLocationInTile38(keyLocation interface{}) (result bool, err error)
	GetLocationInTile38(keyLocation interface{}) (result []float64, err error)
	GetLocationItemNearby(lat, long, radius float64, fields map[string][2]interface{}, pageNumber, pageSize int64) (result map[string]float64, err error)
	SetLocationItemToTile38(keyLocation interface{}, lat, lng float64, fields map[string]interface{}) (err error)
}
