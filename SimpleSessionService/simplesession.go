package SimpleSessionService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/SimpleSessionService/simplesession/thrift/gen-go/simplesession"
	"github.com/OpenStars/EtcdBackendService/SimpleSessionService/simplesession/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type simpleSessionClient struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *simpleSessionClient) GetSession(sskey string) (*simplesession.TUserSessionInfo, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).GetSession(context.Background(), simplesession.TSessionKey(sskey))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r == nil || r.Error == simplesession.TErrorCode_EFailed || r.UserInfo == nil {
		return nil, errors.New("Backedn services:" + m.sid + " err:" + r.Error.String())
	}
	return r.UserInfo, nil
}

func (m *simpleSessionClient) CreateSession(uid int64) (string, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	userinfo := &simplesession.TUserSessionInfo{
		Version: 1,
		Code:    simplesession.TSessionCode_EFullRight,
		UID:     simplesession.TUID(uid),
	}
	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).CreateSession(context.Background(), userinfo)
	if err != nil {
		return "", errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Error == simplesession.TErrorCode_EFailed || r == nil || r.Session == nil {
		return "", errors.New("Backedn services:" + m.sid + " err:" + r.Error.String())
	}
	return string(*r.Session), nil
}

func (m *simpleSessionClient) RemoveSession(sskey string) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).RemoveSession(context.Background(), simplesession.TSessionKey(sskey))
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r == false {
		return errors.New("Backedn services:" + m.sid + " err: remove session false")
	}
	return nil
}

func (m *simpleSessionClient) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
