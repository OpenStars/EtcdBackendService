package vtpnotifystorage

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewVTPNotifyStorageService(etcdServers []string, serviceID, defaulHost, defaultPort string) VTPNotifyStorageService {
	// return sv
	notifystorageService := &vtpnotifystorage{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if notifystorageService.etcdManager == nil {
		return notifystorageService
	}
	err := notifystorageService.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}

	return notifystorageService
}
