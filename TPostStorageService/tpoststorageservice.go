package TPostStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
	"github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tpoststorageservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *tpoststorageservice) GetData(key int64) (*TPostStorageService.TPostItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).GetData(context.Background(), TPostStorageService.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TPostStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *tpoststorageservice) PutData(key int64, data *TPostStorageService.TPostItem) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).PutData(context.Background(), TPostStorageService.TKey(key), data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return nil
}
func (m *tpoststorageservice) RemoveData(key int64) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	_, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).RemoveData(context.Background(), TPostStorageService.TKey(key))
	defer client.BackToPool()
	return err
}
func (m *tpoststorageservice) GetListDatas(listkey []int64) ([]*TPostStorageService.TPostItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	var tlistkeys []TPostStorageService.TKey
	for i := 0; i < len(listkey); i++ {
		tlistkeys = append(tlistkeys, TPostStorageService.TKey(listkey[i]))
	}
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	r, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).GetListDatas(context.Background(), tlistkeys)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.ListDatas == nil || len(r.ListDatas) == 0 {
		return nil, errors.New("Backend service:" + m.sid + " list key not found")
	}
	if r.ErrorCode != TPostStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.ListDatas, nil
}
func (m *tpoststorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
