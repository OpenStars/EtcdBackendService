package StringMapKV

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/StringMapKV/stringmapkv/thrift/gen-go/OpenStars/Common/StringMapKV"
	"github.com/OpenStars/EtcdBackendService/StringMapKV/stringmapkv/transports"
	"github.com/OpenStars/GoEndpointManager"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type stringMapKV struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *stringMapKV) GetData(key string) (string, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetStringMapKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*StringMapKV.StringMapKVServiceClient).GetData(context.Background(), StringMapKV.TKey(key))
	if err != nil {
		return "", errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Data == nil || r.Data.Value == "" {
		return "", errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != StringMapKV.TErrorCode_EGood {
		return "", errors.New("Backedn services:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}
func (m *stringMapKV) PutData(key, value string) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetStringMapKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*StringMapKV.StringMapKVServiceClient).PutData(context.Background(), StringMapKV.TKey(key), &StringMapKV.TStringValue{Value: value})
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return nil
}

func (m *stringMapKV) DeleteKey(key string) error {
	return nil
	// client := transports.GetStringMapKVServiceCompactClient(m.host, m.port)
	// if client == nil || client.Client == nil {
	// 	return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	// }

	// _, err := client.Client.(*StringMapKV.StringMapKVServiceClient).DeleteData(context.Background(), StringMapKV.TKey(key))
	// defer client.BackToPool()
	// return err
}

func (m *stringMapKV) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
