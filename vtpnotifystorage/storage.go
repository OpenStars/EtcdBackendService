package vtpcalllogservice

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpnotifystorage/thrift/gen-go/OpenStars/notifystorage"
	"github.com/OpenStars/EtcdBackendService/vtpnotifystorage/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type vtpnotifystorage struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *vtpnotifystorage) GetData(key int64) (*notifystorage.TNotifyStorage, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetNotifyStorageVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*notifystorage.TNotifyStorageServiceClient).GetData(context.Background(), notifystorage.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != notifystorage.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}

	if r.Data.ID <= 0 {
		return nil, errors.New("Data not existed")
	}

	return r.Data, nil
}

func (m *vtpnotifystorage) GetMultiData(keys []notifystorage.TKey) (map[notifystorage.TKey]*notifystorage.TNotifyStorage, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetNotifyStorageVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*notifystorage.TNotifyStorageServiceClient).GetMultiData(context.Background(), keys)

	if err != nil || r == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *vtpnotifystorage) PutData(id int64, data *notifystorage.TNotifyStorage) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetNotifyStorageVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*notifystorage.TNotifyStorageServiceClient).PutData(context.Background(), notifystorage.TKey(id), data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r != notifystorage.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

func (m *vtpnotifystorage) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
