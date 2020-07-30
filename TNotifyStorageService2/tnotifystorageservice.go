package TNotifyStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/TNotifyStorageService2/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
	"github.com/OpenStars/EtcdBackendService/TNotifyStorageService2/tnotifystorageservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tnotifytorageservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *tnotifytorageservice) GetData(ID int64) (*TNotifyStorageService.TNotifyItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetData(context.Background(), ID)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TNotifyStorageService.TErrorCode_EGood || r.Data == nil || r.Data.ID == 0 {
		return nil, nil
	}
	return r.Data, nil
}

func (m *tnotifytorageservice) PutData(ID int64, data *TNotifyStorageService.TNotifyItem) (bool, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	rs, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).PutData(context.Background(), ID, data)
	if err != nil {
		return false, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if rs != TNotifyStorageService.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *tnotifytorageservice) RemoveData(ID int64) (bool, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	rs, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).RemoveData(context.Background(), ID)
	if err != nil {
		return false, err
	}
	defer client.BackToPool()
	return rs, nil
}

func (m *tnotifytorageservice) GetListDatas(listkey []int64) ([]*TNotifyStorageService.TNotifyItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	var tlistkeys []TNotifyStorageService.TKey
	for i := 0; i < len(listkey); i++ {
		tlistkeys = append(tlistkeys, TNotifyStorageService.TKey(listkey[i]))
	}
	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	r, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetListData(context.Background(), listkey)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	if r.Datass == nil || len(r.Datass) == 0 {
		return nil, errors.New("Backend service:" + m.sid + " list key not found")
	}
	if r.ErrorCode != TNotifyStorageService.TErrorCode_EGood || r.Datass == nil || len(r.Datass) == 0 {
		return nil, nil
	}
	var checklist []*TNotifyStorageService.TNotifyItem
	for _, item := range r.Datass {
		if item.ID != 0 {
			checklist = append(checklist, item)
		}
	}
	return checklist, nil
}
func (m *tnotifytorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
