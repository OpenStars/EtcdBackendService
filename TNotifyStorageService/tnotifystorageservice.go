package TNotifyStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/tnotifystorageservice/thrift/gen-go/OpenStars/Common/TNotifyStorageService"
	"github.com/OpenStars/backendclients/go/tnotifystorageservice/transports"
	"github.com/OpenStars/backendclients/go/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tnotifytorageservice struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *tnotifytorageservice) GetData(key int64) (*TNotifyStorageService.TNotifyItem, error) {
	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetData(context.Background(), key)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TNotifyStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *tnotifytorageservice) PutData(key int64, data *TNotifyStorageService.TNotifyItem) error {
	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).PutData(context.Background(), key, data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return nil
}

func (m *tnotifytorageservice) RemoveData(key int64) error {
	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	_, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).RemoveData(context.Background(), key)
	defer client.BackToPool()
	return err
}

func (m *tnotifytorageservice) GetListDatas(listkey []int64) ([]*TNotifyStorageService.TNotifyItem, error) {
	var tlistkeys []TPostStorageService.TKey
	for i := 0; i < len(listkey); i++ {
		tlistkeys = append(tlistkeys, TPostStorageService.TKey(listkey[i]))
	}
	client := transports.GetTNotifyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	r, err := client.Client.(*TNotifyStorageService.TNotifyStorageServiceClient).GetListData(context.Background(), listkey)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	if r.Datass == nil || len(r.Datass) == 0 {
		return nil, errors.New("Backend service:" + m.sid + " list key not found")
	}
	if r.ErrorCode != TNotifyStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Datass, nil
}
func (m *tnotifytorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
