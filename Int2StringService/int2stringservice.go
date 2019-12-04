package Int2StringService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/i2skv/thrift/gen-go/OpenStars/Common/I2SKV"
	"github.com/OpenStars/backendclients/go/i2skv/transports"
)

type Int2StringService struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *Int2StringService) PutData(key int64, value string) error {
	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	tvalue := &I2SKV.TStringValue{
		Value: value,
	}
	_, err := client.Client.(*I2SKV.TI2StringServiceClient).PutData(context.Background(), tkey, tvalue)
	defer client.BackToPool()
	if err != nil {
		return errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	return nil

}

func (m *Int2StringService) GetData(key int64) (string, error) {
	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	r, err := client.Client.(*I2SKV.TI2StringServiceClient).GetData(context.Background(), tkey)
	if err != nil {
		return "", errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r.Data == nil || r.ErrorCode != I2SKV.TErrorCode_EGood || r.Data.Value == "" {
		return "", errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *Int2StringService) handleEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewInt2StringService(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) Int2StringServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		// log.Println("Load endpoit ", serviceID, "err", err.Error())
		log.Println("Init Local Int2StringService sid:", defaultEndpoint.ServiceID, "host:", defaultEndpoint.Host, "port:", defaultEndpoint.Port)
		return &Int2StringService{
			host: defaultEndpoint.Host,
			port: defaultEndpoint.Port,
			sid:  defaultEndpoint.ServiceID,
		}
	}
	sv := &Int2StringService{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handleEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd Int2StringService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
