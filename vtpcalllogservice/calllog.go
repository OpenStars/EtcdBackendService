package vtpuserinfoservice

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpcalllogservice/thrift/gen-go/OpenStars/calllog"
	"github.com/OpenStars/EtcdBackendService/vtpcalllogservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type vtpcalllogservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *vtpcalllogservice) GetData(key string) (*calllog.TCallLog, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCallLogVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*calllog.TCallLogServiceClient).GetData(context.Background(), calllog.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != calllog.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}

	if r.Data.CallType == "" && r.Data.CallPhone == "" {
		return nil, errors.New("Data not existed")
	}

	return r.Data, nil
}

func (m *vtpcalllogservice) GetMultiData(keys []calllog.TKey) (map[calllog.TKey]*calllog.TCallLog, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCallLogVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*calllog.TCallLogServiceClient).GetMultiData(context.Background(), keys)

	if err != nil || r == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *vtpcalllogservice) PutData(lognumber string, data *calllog.TCallLog) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCallLogVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*calllog.TCallLogServiceClient).PutData(context.Background(), calllog.TKey(lognumber), data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r != calllog.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

func (m *vtpcalllogservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
