package vtpuserinfoservice

import "github.com/OpenStars/EtcdBackendService/vtpuserinfoservice/thrift/gen-go/openstars/enduservtp"

type VTPEndUserService interface {
	GetData(key int64) (*enduservtp.TEndUserVTP, error)
	GetMultiData(keys []enduservtp.TKey) (map[enduservtp.TKey]*enduservtp.TEndUserVTP, error)
	PutData(uid int64, data *enduservtp.TEndUserVTP) error
}
