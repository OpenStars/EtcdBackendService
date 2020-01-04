package TMediaStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/tmediastorageservice/thrift/gen-go/OpenStars/Common/TMediaStorageService"
	"github.com/OpenStars/backendclients/go/tmediastorageservice/transports"

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

	r, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).GetData(context.Background(), TMediaStorageService.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
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

	_, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).PutData(context.Background(), TMediaStorageService.TKey(key), data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return nil
}

func (m *tmediastorageservice) RemoveData(key int64) error {
	client := transports.GetTMediaStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).RemoveData(context.Background(), TMediaStorageService.TKey(key))
	defer client.BackToPool()
	return err
}

func (m *tmediastorageservice) GetListData(listkey []int64) (r []*TMediaStorageService.TMediaItem, err error) {
	client := transports.GetTMediaStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	var listTkey []TMediaStorageService.TKey
	for _, key := range listkey {
		listTkey = append(listTkey, TMediaStorageService.TKey(key))
	}
	rs, err := client.Client.(*TMediaStorageService.TMediaStorageServiceClient).GetListData(context.Background(), listTkey)

	if err != nil || rs == nil || len(rs.ListDatas) == 0 {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	return rs.ListDatas, nil
}

func (m *tmediastorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
