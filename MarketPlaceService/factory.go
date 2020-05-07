package MarketPlaceService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewTMarketPlaceItemService(etcdServers []string, serviceID, defaulHost, defaultPort string) TMarketPlaceItemService {
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
	marketplace := &marketplaceitemservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if marketplace.etcdManager == nil {
		return marketplace
	}
	err := marketplace.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// postsv.etcdManager.GetAllEndpoint(serviceID)
	return marketplace
}
