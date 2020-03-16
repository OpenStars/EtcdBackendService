package TMediaStorageService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewTMediaStorageService(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) TMediaStorageServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local TMediaStorageService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &tmediastorageservice{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &tmediastorageservice{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd TPostStorageService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}

func NewTMediaStorageService2(serviceID string, etcdServers []string, defaultHost string, defaultPort string) TMediaStorageServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local TMediaStorageService sid:", serviceID, "host:", defaultHost, "port:", defaultPort)
		return &tmediastorageservice{
			host: defaultHost,
			port: defaultPort,
			sid:  serviceID,
		}
	}
	sv := &tmediastorageservice{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd TPostStorageService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
