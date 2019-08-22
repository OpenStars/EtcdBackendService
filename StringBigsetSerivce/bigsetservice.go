package StringBigsetSerivce

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/backendclients/go/bigset/transports"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type StringBigsetService struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *StringBigsetService) BsPutItem(bskey generic.TStringKey, item *generic.TItem) error {

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(context.Background(), bskey, item)

	if err != nil {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	if r.Error != generic.TErrorCode_EGood {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + r.Error.String())
	}
	return nil

}

func (m *StringBigsetService) BsGetItem(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error) {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(context.Background(), bskey, itemkey)
	if err != nil {
		return nil, err
	}

	if r.Error != generic.TErrorCode_EGood {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + r.Error.String())
	}
	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount(bskey generic.TStringKey) (int64, error) {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(context.Background(), bskey)

	if err != nil {
		return -1, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	if r <= 0 {
		return -1, errors.New("StringBigsetSerice: " + m.sid + " bigset key don't have any item")
	}
	return r, nil

}

func (m *StringBigsetService) BsGetSlice(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(context.Background(), bskey, fromPos, count)
	if err != nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || len(rs.Items.Items) == 0 {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsMultiPut(bskey generic.TStringKey, lsItems []*generic.TItem) error {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(context.Background(), bskey, itemset, false, false)
	if err != nil {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	if rs.Error != generic.TErrorCode_EGood {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	return nil
}

func (m *StringBigsetService) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewStringBigsetServiceModel(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) StringBigsetServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		// log.Println("Load endpoit ", serviceID, "err", err.Error())
		log.Println("Init Local StringBigsetSerivce sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &StringBigsetService{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &StringBigsetService{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd StringBigsetSerivce sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
