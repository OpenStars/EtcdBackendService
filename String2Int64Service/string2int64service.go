package String2Int64Service

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/backendclients/go/s2i64kv/thrift/gen-go/OpenStars/Common/S2I64KV"
	"github.com/OpenStars/backendclients/go/s2i64kv/transports"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type String2Int64Service struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *String2Int64Service) PutData(key string, value int64) error {
	client := transports.GetS2I64CompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to model")
	}

	tkey := S2I64KV.TKey(key)
	tvalue := &S2I64KV.TI64Value{
		Value: value,
	}

	_, err := client.Client.(*S2I64KV.TString2I64KVServiceClient).PutData(context.Background(), tkey, tvalue)
	defer client.BackToPool()
	if err != nil {
		return errors.New("String2Int64Service sid: " + m.sid + " address: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	return nil
}

func (m *String2Int64Service) GetData(key string) (int64, error) {
	client := transports.GetS2I64CompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to model")
	}

	tkey := S2I64KV.TKey(key)
	r, err := client.Client.(*S2I64KV.TString2I64KVServiceClient).GetData(context.Background(), tkey)
	if err != nil {
		return -1, errors.New("String2Int64Service sid: " + m.sid + " address: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r.Data == nil || r.ErrorCode != S2I64KV.TErrorCode_EGood || r.Data.Value <= 0 {
		return -1, errors.New("Can not found key")
	}
	return r.Data.Value, nil
}

func (m *String2Int64Service) handleEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewString2Int64Service(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) String2Int64ServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		// log.Println("Load endpoit ", serviceID, "err", err.Error())
		log.Println("Init Local String2Int64Service sid:", defaultEndpoint.ServiceID, "host:", defaultEndpoint.Host, "port:", defaultEndpoint.Port)
		return &String2Int64Service{
			host: defaultEndpoint.Host,
			port: defaultEndpoint.Port,
			sid:  defaultEndpoint.ServiceID,
		}
	}
	sv := &String2Int64Service{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handleEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd String2Int64Service sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
