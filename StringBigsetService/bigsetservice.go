package StringBigsetService

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var reconnect = true
var mureconnect sync.Mutex

//go:generate easytags $GOFILE json,xml

type StringBigsetService struct {
	host             string
	port             string
	sid              string
	epm              GoEndpointBackendManager.EndPointManagerIf
	etcdManager      *GoEndpointManager.EtcdBackendEndpointManager
	db               *sql.DB
	isSaveDataBackup bool
	isGetDataBackup  bool
	standardSid      string

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

type MySqlConfig struct {
	Protocol string `json:"protocol" xml:"protocol"`
	UserName string `json:"user_name" xml:"user_name"`
	Password string `json:"password" xml:"password"`
	Schema   string `json:"schema" xml:"schema"`
	IdleTime string `json:"idle_time" xml:"idle_time"`
	LifeTime string `json:"life_time" xml:"life_time"`
	Idle     string `json:"idle" xml:"idle"`
	OpenConn string `json:"open_conn" xml:"open_conn"`
	Host     string `json:"host" xml:"host"`
	Port     string `json:"port" xml:"port"`
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
		go m.notifyEndpointError()
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err = client.Client.(*generic.TStringBigSetKVServiceClient).TotalStringKeyCount(ctx)

	if err != nil {
		go m.notifyEndpointError()
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetListKey(ctx, fromIndex, count)

	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isSaveDataBackup {
		go m.PutToBackupDB(string(bskey), string(item.Key), string(item.Value))
	}

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	// log.Println("BsPutItem host", m.host+":"+m.port)
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(ctx, bskey, item)

	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isGetDataBackup {
		return m.BsRangeQueryBackupDB(string(bskey), startKey, endKey)
	}

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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || len(rs.Items.Items) == 0 {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: ")
	}
	return rs.Items.Items, nil
}

// BsRangeQueryByPage get >= startkey && <= endkey có chia page theo begin and end
func (m *StringBigsetService) BsRangeQueryByPage(bskey generic.TStringKey, startKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error) {
	if m.db != nil && m.isGetDataBackup {
		return m.GetRangeQueryByPageBackupDB(string(bskey), startKey, endKey, begin, end)
	}

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
		go m.notifyEndpointError()
		return nil, -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isGetDataBackup {
		return m.GetItemBackupDB(string(bskey), string(itemkey))
	}

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
	// fmt.Printf("[BsGetItem] get client host = %s, %s, key = %s, %s \n", m.host, m.port, bskey, itemkey)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(ctx, bskey, itemkey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	if r == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: r = nil")
	}
	if r.Error != generic.TErrorCode_EGood || r.Item == nil || r.Item.Key == nil {
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + r.Error.String())
	}

	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount(bskey generic.TStringKey) (int64, error) {
	if m.db != nil && m.isGetDataBackup {
		return m.getTotalCountFromBackupDB(bskey)
	}

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
		go m.notifyEndpointError()
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(ctx, bskey)

	if err != nil {
		go m.notifyEndpointError()
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetBigSetInfoByName(ctx, bskey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, nil
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
	_, err := client.Client.(*generic.TStringBigSetKVServiceClient).RemoveAll(ctx, bskey)
	if err != nil {
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
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
		return nil, nil
	}
	return rs.Info, nil
}

func (m *StringBigsetService) BsGetSlice(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	if m.db != nil && m.isGetDataBackup {
		return m.BsGetSliceBackupDB(bskey, fromPos, count)
	}

	if count == 0 {
		return nil, errors.New("Empty data")
	}
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(ctx, bskey, fromPos, count)
	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isGetDataBackup {
		return m.BsGetSliceRBackupDB(bskey, fromPos, count)
	}

	if count == 0 {
		return nil, errors.New("Empty data")
	}
	// log.Println("host", m.host, "port", m.port, "bskey", string(bskey))
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceR(ctx, bskey, fromPos, count)
	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil {
		go m.RemoveItemBackupDB(string(bskey), string(itemkey))
	}

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
		go m.notifyEndpointError()
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRemoveItem(ctx, bskey, itemkey)
	if err != nil {
		go m.notifyEndpointError()
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
	// todo bsmultiput
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
		go m.notifyEndpointError()
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(ctx, bskey, itemset, false, false)
	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isGetDataBackup {
		return m.BsGetSliceFromItemRBackupDB(bskey, fromKey, count)
	}

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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItem(ctx, bskey, fromKey, count)
	if err != nil {
		go m.notifyEndpointError()
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
	if m.db != nil && m.isGetDataBackup {
		return m.BsGetSliceFromItemRBackupDB(bskey, fromKey, count)
	}

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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItemR(ctx, bskey, fromKey, count)
	if err != nil {
		go m.notifyEndpointError()
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

	log.Println("Init StringBigset Service sid", serviceID, "address", defaultEnpoint.Host+":"+defaultEnpoint.Port)

	stringbs := &StringBigsetService{
		host:        defaultEnpoint.Host,
		port:        defaultEnpoint.Port,
		sid:         defaultEnpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
		bot_chatID:  0,
		bot_token:   "",
		botClient:   nil,
	}
	bot, err := tgbotapi.NewBotAPI(stringbs.bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	stringbs.botClient = nil
	if stringbs.etcdManager == nil {
		return stringbs
	}
	err = stringbs.etcdManager.SetDefaultEntpoint(serviceID, defaultEnpoint.Host, defaultEnpoint.Port)
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

	log.Println("Init StringBigset Service sid", sid, "address", defaultEndpointsHost+":"+defaultEndpointPort)
	stringbs := &StringBigsetService{
		host:        defaultEndpointsHost,
		port:        defaultEndpointPort,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
		bot_chatID:  -1001469468779,
		bot_token:   "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:   nil,
	}
	bot, err := tgbotapi.NewBotAPI(stringbs.bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	stringbs.botClient = nil
	if stringbs.etcdManager == nil {
		return stringbs
	}
	err = stringbs.etcdManager.SetDefaultEntpoint(sid, defaultEndpointsHost, defaultEndpointPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return nil
	}
	return stringbs
}

func NewClient(etcdEndpoints []string, sid string, defaultEndpointsHost string, defaultEndpointPort string) Client {

	log.Println("Init StringBigset Service sid", sid, "address", defaultEndpointsHost+":"+defaultEndpointPort)
	stringbs := &StringBigsetService{
		host:        defaultEndpointsHost,
		port:        defaultEndpointPort,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
		bot_chatID:  -1001469468779,
		bot_token:   "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:   nil,
	}
	bot, err := tgbotapi.NewBotAPI(stringbs.bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	stringbs.botClient = nil
	if stringbs.etcdManager == nil {
		return stringbs
	}
	err = stringbs.etcdManager.SetDefaultEntpoint(sid, defaultEndpointsHost, defaultEndpointPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return nil
	}
	return stringbs
}

func NewClientWithMonitor(etcdEndpoints []string, sid string, host string, port string, bot_token string, bot_chatID int64) Client {
	// 1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg
	// -1001469468779
	log.Println("Init StringBigset Service sid", sid, "address", host+":"+port)
	stringbs := &StringBigsetService{
		host:        host,
		port:        port,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
		botClient:   nil,
		bot_chatID:  bot_chatID,
		bot_token:   bot_token,
	}
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	if stringbs.etcdManager == nil {
		return stringbs
	}
	err = stringbs.etcdManager.SetDefaultEntpoint(sid, host, port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return nil
	}
	return stringbs
}

// ================================================== Version 2 ===============================================================

func (m *StringBigsetService) notifyEndpointError() {
	if m.botClient != nil {

		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện endpoint sid "+m.sid+" address "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}

}

func (m *StringBigsetService) TotalStringKeyCount2() (r int64, err error) {

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
		go m.notifyEndpointError()
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err = client.Client.(*generic.TStringBigSetKVServiceClient).TotalStringKeyCount(ctx)

	if err != nil {
		go m.notifyEndpointError()
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	return r, nil
}

func (m *StringBigsetService) GetListKey2(fromIndex int64, count int32) ([]string, error) {

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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetListKey(ctx, fromIndex, count)

	if err != nil {
		go m.notifyEndpointError()
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

func (m *StringBigsetService) BsPutItem2(bskey generic.TStringKey, item *generic.TItem) (bool, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	// log.Println("BsPutItem host", m.host+":"+m.port)
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(ctx, bskey, item)

	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	return true, nil

}

func (m *StringBigsetService) BsRangeQuery2(bskey generic.TStringKey, startKey generic.TItemKey, endKey generic.TItemKey) ([]*generic.TItem, error) {

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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}

// BsRangeQueryByPage get >= startkey && <= endkey có chia page theo begin and end
func (m *StringBigsetService) BsRangeQueryByPage2(bskey generic.TStringKey, startKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error) {
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
		go m.notifyEndpointError()
		return nil, -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, bskey, startKey, endKey)
	if err != nil {
		go m.notifyEndpointError()
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

func (m *StringBigsetService) BsGetItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (*generic.TItem, error) {
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
	// fmt.Printf("[BsGetItem] get client host = %s, %s, key = %s, %s \n", m.host, m.port, bskey, itemkey)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(ctx, bskey, itemkey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	if r.Error != generic.TErrorCode_EGood || r.Item == nil || r.Item.Key == nil {
		return nil, nil
	}
	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount2(bskey generic.TStringKey) (int64, error) {
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
		go m.notifyEndpointError()
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(ctx, bskey)

	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r <= 0 {
		return 0, nil
	}
	return r, nil

}

func (m *StringBigsetService) GetBigSetInfoByName2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error) {
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
		go m.notifyEndpointError()
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
		return nil, nil
	}
	return rs.Info, nil

}

func (m *StringBigsetService) RemoveAll2(bskey generic.TStringKey) (bool, error) {
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
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Client.(*generic.TStringBigSetKVServiceClient).RemoveAll(ctx, bskey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	return true, nil
}
func (m *StringBigsetService) CreateStringBigSet2(bskey generic.TStringKey) (*generic.TStringBigSetInfo, error) {
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).CreateStringBigSet(ctx, bskey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, nil
	}
	return rs.Info, nil
}

func (m *StringBigsetService) BsGetSlice2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	if count == 0 {
		return nil, nil
	}
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(ctx, bskey, fromPos, count)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceR2(bskey generic.TStringKey, fromPos int32, count int32) ([]*generic.TItem, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	// log.Println("host", m.host, "port", m.port, "bskey", string(bskey))
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceR(ctx, bskey, fromPos, count)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsRemoveItem2(bskey generic.TStringKey, itemkey generic.TItemKey) (bool, error) {
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
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRemoveItem(ctx, bskey, itemkey)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return ok, nil
}

func (m *StringBigsetService) BsMultiPut2(bskey generic.TStringKey, lsItems []*generic.TItem) (bool, error) {
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
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(ctx, bskey, itemset, false, false)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *StringBigsetService) BsGetSliceFromItem2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error) {
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItem(ctx, bskey, fromKey, count)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceFromItemR2(bskey generic.TStringKey, fromKey generic.TItemKey, count int32) ([]*generic.TItem, error) {
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
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItemR(ctx, bskey, fromKey, count)
	if err != nil {
		go m.notifyEndpointError()
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) PutToBackupDB(bsKey, itemKey, value string) {
	_, err := m.db.Exec(fmt.Sprintf("INSERT INTO %s(BsKey, BsItemKey, Val) VALUES(?, ?, ?);", m.standardSid), bsKey, itemKey, value)
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1317")
	}
}

func (m *StringBigsetService) RemoveItemBackupDB(bsKey, itemKey string) {
	_, err := m.db.Exec(fmt.Sprintf("DELETE FROM %s where BsKey = ? and ItemKey = ?;", m.standardSid), bsKey, itemKey)
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1324")
	}
}

func (m *StringBigsetService) GetItemBackupDB(bsKey, itemKey string) (*generic.TItem, error) {
	key := ""
	value := ""

	row := m.db.QueryRow(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey = ?", m.standardSid), bsKey, itemKey)
	if row.Err() != nil {
		err := row.Scan(&key, &value)
		if err != nil {
			log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:276")
		}

		return &generic.TItem{
			Key:   []byte(key),
			Value: []byte(value),
		}, err
	}

	return nil, nil
}

func (m *StringBigsetService) GetRangeQueryByPageBackupDB(bsKey string, startKey, endKey generic.TItemKey, begin, end int64) ([]*generic.TItem, int64, error) {
	totalCount := int64(0)
	row := m.db.QueryRow(fmt.Sprintf("SELECT count(*) FROM %s WHERE BsKey = ? and BsItemKey >= ? and BsItemKey < ?", m.standardSid), bsKey, string(startKey), string(endKey))
	if row.Err() != nil {
		err := row.Scan(&totalCount)
		if err != nil {
			log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1384")
		}
	}

	rows, err := m.db.Query(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey >= ? and BsItemKey < ? limit %d offset %d", m.standardSid, begin, end), bsKey, string(startKey), string(endKey))
	items := make([]*generic.TItem, 0)

	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1397")
		return make([]*generic.TItem, 0), 0, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			items = append(items, item)
		}

		return items, totalCount, nil
	}

	return items, 0, nil
}

func (m *StringBigsetService) getTotalCountFromBackupDB(bskey generic.TStringKey) (int64, error) {
	rs, err := m.db.Query(fmt.Sprintf("SELECT count(*) FROM %s WHERE BsKey = ?", m.standardSid), bskey)
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1425")
		return 0, err
	}

	totalCount := int64(0)

	for rs.Next() {
		err = rs.Scan(&totalCount)
		if err != nil {
			log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1434")
			return 0, err
		}
	}

	return totalCount, nil
}

func (m *StringBigsetService) BsGetSliceFromItemRBackupDB(bskey generic.TStringKey, fromItemKey generic.TItemKey, count int32) (result []*generic.TItem, err error) {
	result = make([]*generic.TItem, 0)

	rows, err := m.db.Query(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey >= ? order by BsItemKey desc limit %d", m.standardSid, count), bskey, string(fromItemKey))
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1398")
		return result, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			result = append(result, item)
		}

		return result, nil
	}

	return result, nil
}

func (m *StringBigsetService) BsGetSliceFromItemBackupDB(bskey generic.TStringKey, fromItemKey generic.TItemKey, count int64) (result []*generic.TItem, err error) {
	result = make([]*generic.TItem, 0)

	rows, err := m.db.Query(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey >= ? order by BsItemKey asc limit %d", m.standardSid, count), bskey, string(fromItemKey))
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1416")
		return result, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			result = append(result, item)
		}

		return result, nil
	}

	return result, nil
}

func (m *StringBigsetService) BsGetSliceRBackupDB(bskey generic.TStringKey, from, count int32) (result []*generic.TItem, err error) {
	rows, err := m.db.Query(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? order by BsItemKey limit %d offset %d", m.standardSid, count, from), bskey)
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1452")
		return result, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			result = append(result, item)
		}

		return result, err
	}

	return result, err
}

func (m *StringBigsetService) BsGetSliceBackupDB(bskey generic.TStringKey, from, count int32) (result []*generic.TItem, err error) {
	rows, err := m.db.Query(fmt.Sprintf("SELECT BsItemKey, Val FROM %s WHERE BsKey = ? limit %d offset %d", m.standardSid, count, from), bskey)
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1452")
		return result, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			result = append(result, item)
		}

		return result, err
	}

	return result, err
}

func (m *StringBigsetService) BsRangeQueryBackupDB(bsKey string, begin generic.TItemKey, end generic.TItemKey) (result []*generic.TItem, err error) {
	rows, err := m.db.Query(fmt.Sprintf("SELECT BsKey, BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey >= ? and BsItemKey < ?", m.standardSid), bsKey, string(begin), string(end))
	if err != nil {
		log.Println(err.Error(), "err.Error() StringBigsetService/bigsetservice.go:1452")
		return result, err
	}

	item := &generic.TItem{}
	itemKey := ""
	value := ""

	if rows != nil {
		for rows.Next() {
			err := rows.Scan(itemKey, value)
			if err != nil {
				log.Fatal(err)
			}
			item.Key = []byte(itemKey)
			item.Value = []byte(value)

			result = append(result, item)
		}

		return result, err
	}

	return result, err
}

func NewClientSyncTiKv(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint, cfg MySqlConfig, isSaveDataBackup, isGetDataBackup bool) StringBigsetServiceIf {
	log.Println("Init StringBigset Service sid", serviceID, "address", defaultEnpoint.Host+":"+defaultEnpoint.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s?collation=utf8_bin&interpolateParams=true",
		cfg.UserName, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.Schema))
	if err != nil {
		log.Println(err.Error(), "err.Error() can't connect to tikv StringBigsetService/bigsetservice.go:1438")
		return nil
	}

	stringbs := &StringBigsetService{
		host:             defaultEnpoint.Host,
		port:             defaultEnpoint.Port,
		sid:              defaultEnpoint.ServiceID,
		isSaveDataBackup: isSaveDataBackup,
		isGetDataBackup:  isGetDataBackup,
		db:               db,
		etcdManager:      GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
		bot_chatID:       0,
		bot_token:        "",
		botClient:        nil,
	}

	bot, err := tgbotapi.NewBotAPI(stringbs.bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	stringbs.botClient = nil
	if stringbs.etcdManager == nil {
		return stringbs
	}
	err = stringbs.etcdManager.SetDefaultEntpoint(serviceID, defaultEnpoint.Host, defaultEnpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}

	standardSid := strings.ReplaceAll(strings.ReplaceAll(serviceID, "/", "_"), "-", "_")
	_, err = db.Exec(fmt.Sprintf(`create table if not exists %s
		(
			BsKey                varchar(255),
			BsItemKey            varchar(255),
			Val            	 text,
			primary key (BsKey, BsItemKey)
		); `, standardSid))
	stringbs.standardSid = standardSid
	if err != nil {
		log.Println(err.Error(), "err.Error() create table failed StringBigsetService/bigsetservice.go:691")
	}

	return stringbs
}
