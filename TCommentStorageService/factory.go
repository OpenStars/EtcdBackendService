package TCommentStorageService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewTCommentStorageService(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) TCommentStorageServiceIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local TCommentStorageService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &tcommentstorageservice{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &tcommentstorageservice{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd TPostStorageService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
