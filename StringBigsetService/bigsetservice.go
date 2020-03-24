package StringBigsetService

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/OpenStars/GoEndpointManager"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/transports"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var reconnect = true
var mureconnect sync.Mutex

type StringBigsetService struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *StringBigsetService) TotalStringKeyCount() (r int64, err error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err = client.Client.(*generic.TStringBigSetKVServiceClient).TotalStringKeyCount(ctx)

	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	return r, nil
}

func (m *StringBigsetService) GetListKey(fromIndex int64, count int32) ([]string, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetListKey(ctx, fromIndex, count)

	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	var listKey []string
	for _, item := range r {
		listKey = append(listKey, string(item))
	}
	return listKey, nil
}

func (m *StringBigsetService) BsPutItem(bskey generic.TStringKey, item *generic.TItem) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(ctx, bskey, item)

	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r.Error != generic.TErrorCode_EGood {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + r.Error.String())
	}
	return nil

}

func (m *StringBigsetService) BsRangeQuery(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || len(rs.Items.Items) == 0 {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	return rs.Items.Items, nil
}

// BsRangeQueryByPage get >= startkey && <= endkey có chia page theo begin and end
func (m *StringBigsetService) BsRangeQueryByPage(bskey generic.TStringKey, startKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return nil, -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		return nil, -1, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r.Items != nil && r.Items.Items != nil && len(r.Items.Items) > 0 { // pagination
		if begin < 0 {
			begin = 0
		}
		if end > int64(len(r.Items.Items)) {
			end = int64(len(r.Items.Items))
		}
		total := int64(len(r.Items.Items))
		r.Items.Items = r.Items.Items[begin:end]

		return r.Items.Items, total, nil
	}

	return nil, 0, nil
}

func (m *StringBigsetService) BsGetItem(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	fmt.Printf("[BsGetItem] get client host = %s, %s, key = %s, %s \n", m.host, m.port, bskey, itemkey)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(ctx, bskey, itemkey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	if r.Error != generic.TErrorCode_EGood {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + r.Error.String())
	}
	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount(bskey generic.TStringKey) (int64, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(ctx, bskey)

	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r <= 0 {
		return 0, nil
	}
	return r, nil

}

func (m *StringBigsetService) GetBigSetInfoByName(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetBigSetInfoByName(ctx, bskey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, errors.New("Get bigset info by name err " + rs.Error.String())
	}
	return rs.Info, nil

}

func (m *StringBigsetService) RemoveAll(bskey generic.TStringKey) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).RemoveAll(ctx, bskey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	log.Println("BsRemoveAll rs", rs)
	return nil
}
func (m *StringBigsetService) CreateStringBigSet(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).CreateStringBigSet(ctx, bskey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, errors.New("Get bigset info by name err " + rs.Error.String())
	}
	return rs.Info, nil
}

func (m *StringBigsetService) BsGetSlice(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(ctx, bskey, fromPos, count)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}

	if len(rs.Items.Items) == 0 {
		return []*generic.TItem{}, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceR(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceR(ctx, bskey, fromPos, count)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}

	if len(rs.Items.Items) == 0 {
		return []*generic.TItem{}, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsRemoveItem(bskey generic.TStringKey, itemkey generic.TItemKey) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRemoveItem(ctx, bskey, itemkey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if ok == false {
		return errors.New("StringBigsetSerice: " + m.sid + " remove false")
	}
	return nil
}

func (m *StringBigsetService) BsMultiPut(bskey generic.TStringKey, lsItems []*generic.TItem) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(ctx, bskey, itemset, false, false)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood {
		return errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	return nil
}

func (m *StringBigsetService) BsGetSliceFromItem(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItem(ctx, bskey, fromKey, count)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	if len(rs.Items.Items) == 0 {
		return []*generic.TItem{}, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceFromItemR(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItemR(ctx, bskey, fromKey, count)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + rs.Error.String())
	}
	if len(rs.Items.Items) == 0 {
		return []*generic.TItem{}, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewStringBigsetServiceModel(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) StringBigsetServiceIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	stringbs := &StringBigsetService{
		host:        defaultEnpoint.Host,
		port:        defaultEnpoint.Port,
		sid:         defaultEnpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}
	if stringbs.etcdManager == nil {
		return nil
	}
	err := stringbs.etcdManager.SetDefaultEntpoint(serviceID, defaultEnpoint.Host, defaultEnpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// stringbs.etcdManager.GetAllEndpoint(serviceID)
	return stringbs
	// if err != nil {
	// 	log.Println("Init Local StringBigsetSerivce sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &StringBigsetService{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &StringBigsetService{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd StringBigsetSerivce sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv
}

func NewStringBigsetServiceModel2(etcdEndpoints []string, sid string, defaultEndpointsHost string, defaultEndpointPort string) StringBigsetServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdEndpoints, sid)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local StringBigsetSerivce sid:", sid, "host:", defaultEndpointsHost+":"+defaultEndpointPort)
		return &StringBigsetService{
			host: defaultEndpointsHost,
			port: defaultEndpointPort,
			sid:  sid,
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
