package MapPhoneNumber2Pubkey

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/OpenStars/backendclients/go/mapphone2pubkey/thrift/gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV"
	"github.com/OpenStars/backendclients/go/mapphone2pubkey/transports"
)

type MappingPhone2PubkeyServiceModel struct {
	host string
	port string
	sid  string
	epm  GoEndpointBackendManager.EndPointManagerIf
}

func (m *MappingPhone2PubkeyServiceModel) PutData(pubkey string, phonenumber string) error {
	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to model")
	}
	defer client.BackToPool()

	_, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).PutData(context.Background(), pubkey, phonenumber)

	return err
}

func (m *MappingPhone2PubkeyServiceModel) GetPhoneNumberByPubkey(pubkey string) (string, error) {
	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}
	defer client.BackToPool()

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPhoneNumberByPubkey(context.Background(), pubkey)
	if err != nil {
		return "", err
	}
	if r.ErrorCode != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return "", errors.New("Get phonenubmer by pubkey errors " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *MappingPhone2PubkeyServiceModel) GetPubkeyByPhoneNumber(phonenumber string) (string, error) {
	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}
	defer client.BackToPool()

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPubkeyByPhoneNumber(context.Background(), phonenumber)
	if err != nil {
		return "", err
	}
	if r.ErrorCode != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return "", errors.New("Get phonenubmer by pubkey errors " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *MappingPhone2PubkeyServiceModel) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewMappingPhone2Pubkey(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) MappingPhoneNumber2PubkeyServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		// log.Println("Load endpoit ", serviceID, "err", err.Error())
		log.Println("Init Local MappingPhone2Pubkey sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &MappingPhone2PubkeyServiceModel{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &MappingPhone2PubkeyServiceModel{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd MappingPhone2Pubkey sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
