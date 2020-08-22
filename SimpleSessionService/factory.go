package SimpleSessionService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewSimpleSessionClient(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) SimpleSessionClientIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	log.Println("Init Local SimpleSession sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &simpleSessionClient{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &simpleSessionClient{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd SimpleSession sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv

	sessionsv := &simpleSessionClient{
		host:        defaultEndpoint.Host,
		port:        defaultEndpoint.Port,
		sid:         defaultEndpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if sessionsv.etcdManager == nil {
		return nil
	}
	err := sessionsv.etcdManager.SetDefaultEntpoint(serviceID, defaultEndpoint.Host, defaultEndpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// sessionsv.etcdManager.GetAllEndpoint(serviceID)
	return sessionsv

}

func NewSimpleSessionClient2(etcdServers []string, sid, defaultHost, defaultPort string) SimpleSessionClientIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	log.Println("Init Local SimpleSession sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &simpleSessionClient{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &simpleSessionClient{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd SimpleSession sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv

	sessionsv := &simpleSessionClient{
		host:        defaultHost,
		port:        defaultPort,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if sessionsv.etcdManager == nil {
		return sessionsv
	}
	err := sessionsv.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return nil
	}
	// sessionsv.etcdManager.GetAllEndpoint(serviceID)
	return sessionsv

}

func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) Client {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	log.Println("Init Local SimpleSession sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &simpleSessionClient{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &simpleSessionClient{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd SimpleSession sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv

	sessionsv := &simpleSessionClient{
		host:        defaultHost,
		port:        defaultPort,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if sessionsv.etcdManager == nil {
		return sessionsv
	}
	err := sessionsv.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return nil
	}
	// sessionsv.etcdManager.GetAllEndpoint(serviceID)
	return sessionsv

}
