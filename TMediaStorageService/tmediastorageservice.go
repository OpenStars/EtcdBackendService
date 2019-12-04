package TMediaStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/tmediastorageservice/thrift/gen-go/OpenStars/Common/TMediaStorageService"
	"github.com/OpenStars/backendclients/go/tmediastorageservice/transports"
	"github.com/OpenStars/backendclients/go/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tmediastorageservice struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *tmediastorageservice) GetData(key int64) (*TMediaStorageService.TMediaItem, error) {
	client := transports.GetTMediaStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	r, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).GetData(context.Background(), TMediaStorageService.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TMediaStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *tmediastorageservice) PutData(key int64, data *TMediaStorageService.TMediaItem) error {
	client := transports.GetTMediaStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	_, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).PutData(context.Background(), TMediaStorageService.TKey(key), data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	return nil
}

func (m *tmediastorageservice) RemoveData(key int64) error {
	client := transports.GetTMediaStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	_, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).RemoveData(context.Background(), TPostStorageService.TKey(key))
	return err
}
func (m *tmediastorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
