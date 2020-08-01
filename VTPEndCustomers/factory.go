package vtpuserservice

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewVTPEndUserService(etcdServers []string, serviceID, defaulHost, defaultPort string) VTPEndUserService {
	// return sv
	enduser := &vtpenduserservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if enduser.etcdManager == nil {
		return enduser
	}
	err := enduser.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}
	// postsv.etcdManager.GetAllEndpoint(serviceID)
	return enduser
}
