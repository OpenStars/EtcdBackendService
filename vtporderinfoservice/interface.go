package vtporderinfoservice

import "github.com/OpenStars/EtcdBackendService/vtporderinfoservice/thrift/gen-go/OpenStars/orderservice"

type OrderInfoService interface {
	GetData(key int64) (*orderservice.TOrder, error)
	PutData(uid int64, data *orderservice.TOrder) error
}
