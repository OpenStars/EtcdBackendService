package ReportItemService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/TReportItemService/treportitemservice/thrift/gen-go/OpenStars/Platform/MarketPlace"
	"github.com/OpenStars/EtcdBackendService/TReportItemService/treportitemservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type reportitemservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *reportitemservice) GetData(key int64) (*MarketPlace.TReportItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTReportItemServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*MarketPlace.TReportItemServiceClient).GetData(context.Background(), key)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil || r.Data.ID == 0 {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != MarketPlace.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *reportitemservice) PutData(key int64, data *MarketPlace.TReportItem) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTReportItemServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	_, err := client.Client.(*MarketPlace.TReportItemServiceClient).PutData(context.Background(), key, data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	defer client.BackToPool()
	return nil
}

func (m *reportitemservice) RemoveData(key int64) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTReportItemServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	_, err := client.Client.(*MarketPlace.TReportItemServiceClient).RemoveData(context.Background(), key)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}
	defer client.BackToPool()
	return nil
}

func (m *reportitemservice) GetListDatas(listkey []int64) ([]*MarketPlace.TReportItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTReportItemServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}
	r, err := client.Client.(*MarketPlace.TReportItemServiceClient).GetListData(context.Background(), listkey)
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

func (m *reportitemservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
