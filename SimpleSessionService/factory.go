package SimpleSessionService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewSimpleSessionClient(serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) SimpleSessionClientIf {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local SimpleSession sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &simpleSessionClient{
			host: defaultEnpoint.Host,
			port: defaultEnpoint.Port,
			sid:  defaultEnpoint.ServiceID,
		}
	}
	sv := &simpleSessionClient{
		host: ep.Host,
		port: ep.Port,
		sid:  ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd SimpleSession sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	return sv
}
