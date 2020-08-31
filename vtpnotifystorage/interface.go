package vtpcalllogservice

import "github.com/OpenStars/EtcdBackendService/vtpnotifystorage/thrift/gen-go/OpenStars/notifystorage"

type VTPNotifyStorageService interface {
	GetData(key int64) (*notifystorage.TNotifyStorage, error)
	GetMultiData(keys []notifystorage.TKey) (map[notifystorage.TKey]*notifystorage.TNotifyStorage, error)
	PutData(id int64, data *notifystorage.TNotifyStorage) error
}
