package TCommentStorageService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/tcommentstorageservice/thrift/gen-go/OpenStars/Common/TCommentStorageService"
	"github.com/OpenStars/backendclients/go/tcommentstorageservice/transports"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type tcommentstorageservice struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *tcommentstorageservice) GetData(key int64) (*TCommentStorageService.TCommentItem, error) {
	client := transports.GetTCommentStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*TCommentStorageService.TCommentStorageServiceClient).GetData(context.Background(), TCommentStorageService.TKey(key))
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != TCommentStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *tcommentstorageservice) PutData(key int64, data *TCommentStorageService.TCommentItem) error {
	client := transports.GetTCommentStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*TCommentStorageService.TCommentStorageServiceClient).PutData(context.Background(), TCommentStorageService.TKey(key), data)
	if err != nil {
		return errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return nil
}

func (m *tcommentstorageservice) RemoveData(key int64) error {
	client := transports.GetTCommentStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	_, err := client.Client.(*TCommentStorageService.TCommentStorageServiceClient).RemoveData(context.Background(), TCommentStorageService.TKey(key))
	defer client.BackToPool()
	return err
}
func (m *tcommentstorageservice) GetListDatas(listkey []int64) ([]*TCommentStorageService.TCommentItem, error) {
	var tlistkeys []TCommentStorageService.TKey
	for i := 0; i < len(listkey); i++ {
		tlistkeys = append(tlistkeys, TCommentStorageService.TKey(listkey[i]))
	}
	client := transports.GetTCommentStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*TCommentStorageService.TCommentStorageServiceClient).GetListData(context.Background(), tlistkeys)
	if err != nil {
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.ListDatas == nil || len(r.ListDatas) == 0 {
		return nil, errors.New("Backend service:" + m.sid + " list key not found")
	}
	if r.ErrorCode != TCommentStorageService.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.ListDatas, nil
}
func (m *tcommentstorageservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}
