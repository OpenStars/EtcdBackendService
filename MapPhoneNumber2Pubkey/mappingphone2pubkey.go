package MapPhoneNumber2Pubkey

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/MapPhoneNumber2Pubkey/mapphone2pubkey/thrift/gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV"
	"github.com/OpenStars/EtcdBackendService/MapPhoneNumber2Pubkey/mapphone2pubkey/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type MappingPhone2PubkeyServiceModel struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *MappingPhone2PubkeyServiceModel) PutData(pubkey string, phonenumber string) error {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to model")
	}

	_, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).PutData(context.Background(), pubkey, phonenumber)
	if err != nil {
		return errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
	return err
}

func (m *MappingPhone2PubkeyServiceModel) GetPhoneNumberByPubkey(pubkey string) (string, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPhoneNumberByPubkey(context.Background(), pubkey)
	if err != nil {
		return "", errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.ErrorCode != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return "", errors.New("Get phonenubmer by pubkey errors " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *MappingPhone2PubkeyServiceModel) GetPubkeyByPhoneNumber(phonenumber string) (string, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPubkeyByPhoneNumber(context.Background(), phonenumber)
	if err != nil {
		return "", errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
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

func NewMappingPhone2Pubkey(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) MappingPhoneNumber2PubkeyServiceIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	// log.Println("Load endpoit ", serviceID, "err", err.Error())
	// 	log.Println("Init Local MappingPhone2Pubkey sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &MappingPhone2PubkeyServiceModel{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &MappingPhone2PubkeyServiceModel{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd MappingPhone2Pubkey sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv

	mapphone2pub := &MappingPhone2PubkeyServiceModel{
		host:        defaultEndpoint.Host,
		port:        defaultEndpoint.Port,
		sid:         defaultEndpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if mapphone2pub.etcdManager == nil {
		return nil
	}
	err := mapphone2pub.etcdManager.SetDefaultEntpoint(serviceID, defaultEndpoint.Host, defaultEndpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// mapphone2pub.etcdManager.GetAllEndpoint(serviceID)
	return mapphone2pub

}
