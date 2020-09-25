package vtpcmtorderservice

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

func NewVTPCommentOrderService(etcdServers []string, serviceID, defaulHost, defaultPort string) VTPCommentOrderService {
	// return sv
	cmtorderService := &vtpcommentorderservice{
		host:        defaulHost,
		port:        defaultPort,
		sid:         serviceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if cmtorderService.etcdManager == nil {
		return cmtorderService
	}
	err := cmtorderService.etcdManager.SetDefaultEntpoint(serviceID, defaulHost, defaultPort)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return nil
	}

	return cmtorderService
}
