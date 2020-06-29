package userinfoservice

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewUserInfoService(etcdServers []string, serviceID, defaulHost, defaultPort string) UserInfoService {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	log.Println("Init Local TPostStorageService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &tpoststorageservice{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &tpoststorageservice{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd TPostStorageService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv
	reportitem := &userinforservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if reportitem.etcdManager == nil {
		return reportitem
	}
	err := reportitem.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// postsv.etcdManager.GetAllEndpoint(serviceID)
	return reportitem
}
