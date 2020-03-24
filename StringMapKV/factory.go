package StringMapKV

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewStringMapKV(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) StringMapKVIf {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	log.Println("Init Local StringMapKV sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &stringMapKV{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &stringMapKV{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd StringMapKV sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv

	strinmapkv := &stringMapKV{
		host:        defaultEndpoint.Host,
		port:        defaultEndpoint.Port,
		sid:         defaultEndpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if strinmapkv.etcdManager == nil {
		return nil
	}
	err := strinmapkv.etcdManager.SetDefaultEntpoint(serviceID, defaultEndpoint.Host, defaultEndpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// strinmapkv.etcdManager.GetAllEndpoint(serviceID)
	return strinmapkv
}
