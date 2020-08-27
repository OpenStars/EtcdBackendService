package vtpuserinfoservice

import "github.com/OpenStars/EtcdBackendService/vtpcalllogservice/thrift/gen-go/OpenStars/calllog"

type VTPCallLogService interface {
	GetData(key string) (*calllog.TCallLog, error)
	GetMultiData(keys []calllog.TKey) (map[calllog.TKey]*calllog.TCallLog, error)
	PutData(lognumber string, data *calllog.TCallLog) error
}
