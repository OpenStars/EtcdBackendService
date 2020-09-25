package vtpcmtorderservice

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/vtpcmtorderservice/thrift/gen-go/OpenStars/VTPComment"
	"github.com/OpenStars/EtcdBackendService/vtpcmtorderservice/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type vtpcommentorderservice struct {
	host        string
	port        string
	sid         string
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *vtpcommentorderservice) GetData(key string) (*VTPComment.TVTPComment, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCommentOrderVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*VTPComment.TVTPCommentServiceClient).GetData(context.Background(), VTPComment.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != VTPComment.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}

	if string(r.Data.ID) == "" && r.Data.Text == "" && r.Data.UIDComment == "" {
		return nil, errors.New("Data not existed")
	}

	return r.Data, nil
}

func (m *vtpcommentorderservice) GetMultiData(keys []VTPComment.TKey) (map[VTPComment.TKey]*VTPComment.TVTPComment, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCommentOrderVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*VTPComment.TVTPCommentServiceClient).GetMultiData(context.Background(), keys)

	if err != nil || r == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *vtpcommentorderservice) PutData(lognumber string, data *VTPComment.TVTPComment) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetCommentOrderVTPServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*VTPComment.TVTPCommentServiceClient).PutData(context.Background(), VTPComment.TKey(lognumber), data)
	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r != VTPComment.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

func (m *vtpcommentorderservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
