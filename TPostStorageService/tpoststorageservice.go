package StringMapKV

import (
	"context"
	"errors"

	"github.com/OpenStars/backendclients/go/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
	"github.com/OpenStars/backendclients/go/tpoststorageservice/transports"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tpoststorageservice struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *tpoststorageservice) GetData(key int64) (*TPostStorageService.TPostItem, error) {
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	r, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).GetData(context.Background(), TPostStorageService.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TPostStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *tpoststorageservice) PutData(key int64, data *TPostStorageService.TPostItem) error {
	client := transports.GetTPostStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	_, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).PutData(context.Background(), TPostStorageService.TKey(key), data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	return nil
}

// func (m *stringMapKV) DeleteKey(key string) error {
// 	client := transports.GetStringMapKVServiceCompactClient(m.host, m.port)
// 	if client == nil || client.Client == nil {
// 		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
// 	}
// 	defer client.BackToPool()
// 	_, err := client.Client.(*StringMapKV.StringMapKVServiceClient).DeleteData(context.Background(), StringMapKV.TKey(key))
// 	return err
// }

// func (m *stringMapKV) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
// 	m.host = ep.Host
// 	m.port = ep.Port
// 	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
// }
