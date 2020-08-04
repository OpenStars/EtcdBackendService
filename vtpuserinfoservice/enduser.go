package vtpuserinfoservice

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpuserinfoservice/thrift/gen-go/openstars/enduservtp"
	"github.com/OpenStars/EtcdBackendService/vtpuserinfoservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type vtpenduserservice struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *vtpenduserservice) GetData(key int64) (*enduservtp.TEndUserVTP, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetEndUserVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*enduservtp.TEndUserVTPServiceClient).GetData(context.Background(), enduservtp.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != enduservtp.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}

	if int64(r.Data.UID) == int64(0) && r.Data.PhoneNumber == "" {
		return nil, errors.New("Data not existed")
	}

	return r.Data, nil
}

func (m *vtpenduserservice) GetMultiData(keys []enduservtp.TKey) (map[enduservtp.TKey]*enduservtp.TEndUserVTP, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetEndUserVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*enduservtp.TEndUserVTPServiceClient).GetMultiData(context.Background(), keys)

	if err != nil || r == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *vtpenduserservice) PutData(uid int64, data *enduservtp.TEndUserVTP) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetEndUserVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*enduservtp.TEndUserVTPServiceClient).PutData(context.Background(), enduservtp.TKey(uid), data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r != enduservtp.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

func (m *vtpenduserservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
