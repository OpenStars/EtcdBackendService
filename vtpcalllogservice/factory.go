package vtpcalllogservice

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewVTPCallLogService(etcdServers []string, serviceID, defaulHost, defaultPort string) VTPCallLogService {
	// return sv
	calllogService := &vtpcalllogservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if calllogService.etcdManager == nil {
		return calllogService
	}
	err := calllogService.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}

	return calllogService
}
