package vtporderinfoservice

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewOrderInfoService(etcdServers []string, serviceID, defaulHost, defaultPort string) OrderInfoService {
	// return sv
	orderservice := &orderinfoservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if orderservice.etcdManager == nil {
		return orderservice
	}
	err := orderservice.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}

	return orderservice
}
