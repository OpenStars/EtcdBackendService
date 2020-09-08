package vtporderinfoservice2

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/vtporderinfoservice2/thrift/gen-go/OpenStars/orderservice"
	"github.com/OpenStars/EtcdBackendService/vtporderinfoservice2/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type orderinfoservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *orderinfoservice) GetData(key string) (*orderservice.TOrder, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetOrderInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*orderservice.TOrderServiceClient).GetData(context.Background(), orderservice.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != orderservice.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}

	if int64(r.Data.NgayGiao) == int64(0) && r.Data.MaVandon == "" {
		return nil, errors.New("Data not existed")
	}

	return r.Data, nil
}

func (m *orderinfoservice) PutData(id string, data *orderservice.TOrder) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetOrderInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*orderservice.TOrderServiceClient).PutData(context.Background(), orderservice.TKey(id), data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r != orderservice.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

func (m *orderinfoservice) GetMulti(ids []orderservice.TKey) (map[orderservice.TKey]*orderservice.TOrder, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetOrderInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return map[orderservice.TKey]*orderservice.TOrder{}, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*orderservice.TOrderServiceClient).GetMultiData(context.Background(), ids)
	if err != nil {
		return map[orderservice.TKey]*orderservice.TOrder{}, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()

	return r, nil
}

func (m *orderinfoservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
