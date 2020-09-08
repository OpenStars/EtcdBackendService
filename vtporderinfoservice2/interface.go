package vtporderinfoservice2

import "github.com/OpenStars/EtcdBackendService/vtporderinfoservice2/thrift/gen-go/OpenStars/orderservice"

type OrderInfoService interface {
	GetData(key string) (*orderservice.TOrder, error)
	PutData(uid string, data *orderservice.TOrder) error
	GetMulti(ids []orderservice.TKey) (map[orderservice.TKey]*orderservice.TOrder, error)
}
