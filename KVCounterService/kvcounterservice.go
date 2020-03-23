package KVCounterService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/KVCounterService/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"
	"github.com/OpenStars/EtcdBackendService/KVCounterService/kvcounter/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type KVCounterService struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *KVCounterService) GetValue(genname string) (int64, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetKVCounterCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetValue(context.Background(), genname)
	if err != nil {
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return r, nil

}

func (m *KVCounterService) GetCurrentValue(genname string) (int64, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetKVCounterCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetCurrentValue(context.Background(), genname)
	if err != nil {

		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return r, nil
}

func (m *KVCounterService) GetStepValue(genname string, step int64) (int64, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetKVCounterCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetStepValue(context.Background(), genname, step)
	if err != nil {
		// client = transports.NewGetKVCounterCompactClient(m.host, m.port)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return r, nil
}

func (m *KVCounterService) CreateGenerator(genname string) (int32, error) {

	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetKVCounterCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).CreateGenerator(context.Background(), genname)
	if err != nil {

		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return r, nil

}

func (m *KVCounterService) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewKVCounterServiceModel(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) KVCounterServiceIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	// log.Println("Load endpoit ", serviceID, "err", err.Error())
	// 	log.Println("Init Local KVCounterService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &KVCounterService{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &KVCounterService{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd KVCounterService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv
	kvcounter := &KVCounterService{
		host:        defaultEnpoint.Host,
		port:        defaultEnpoint.Port,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if kvcounter.etcdManager == nil {
		return nil
	}
	err := kvcounter.etcdManager.SetDefaultEntpoint(serviceID, defaultEnpoint.Host, defaultEnpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	kvcounter.etcdManager.GetAllEndpoint(serviceID)
	return kvcounter
}
