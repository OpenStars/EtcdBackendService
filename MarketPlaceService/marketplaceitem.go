package MarketPlaceService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/MarketPlaceService/marketplaceitem/thrift/gen-go/OpenStars/Platform/MarketPlace"
	"github.com/OpenStars/EtcdBackendService/MarketPlaceService/marketplaceitem/transports"

	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type marketplaceitemservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *marketplaceitemservice) GetData(key int64) (*MarketPlace.TMarketPlaceItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTMarketPlaceServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*MarketPlace.TMarketPlaceServiceClient).GetData(context.Background(), key)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != MarketPlace.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *marketplaceitemservice) PutData(key int64, data *MarketPlace.TMarketPlaceItem) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTMarketPlaceServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	_, err := client.Client.(*MarketPlace.TMarketPlaceServiceClient).PutData(context.Background(), key, data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	defer client.BackToPool()
	return nil
}

func (m *marketplaceitemservice) RemoveData(key int64) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTMarketPlaceServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	_, err := client.Client.(*MarketPlace.TMarketPlaceServiceClient).RemoveData(context.Background(), key)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	defer client.BackToPool()
	return nil
}

func (m *marketplaceitemservice) GetListDatas(listkey []int64) ([]*MarketPlace.TMarketPlaceItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTMarketPlaceServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}
	r, err := client.Client.(*MarketPlace.TMarketPlaceServiceClient).GetListData(context.Background(), listkey)
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}
	defer client.BackToPool()
	if r == nil || len(r.Data) == 0 {
		return nil, errors.New("Backend service:" + m.sid + " list key not found")
	}
	if r.ErrorCode != MarketPlace.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *marketplaceitemservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
