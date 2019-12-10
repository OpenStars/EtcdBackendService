package TPostStorageService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewTNotifyStorageService(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) TNotifyStorageServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local TNotifyStorageService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &tnotifytorageservice{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &tnotifytorageservice{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd TNotifyStorageService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
