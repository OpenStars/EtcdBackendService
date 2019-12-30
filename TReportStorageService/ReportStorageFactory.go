/**
 * @author tunghx
 * @email tunghx@sonek.vn
 * @create date 12/14/19 12:18 PM
 * @modify date 12/14/19 12:18 PM
 * @desc [description]
 */

package TReportStorageService

import (
	"log"

	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func NewTReportStorageService(serviceID string, etcdServers []string,
	defaultEnpoint GoEndpointBackendManager.EndPoint,
	bigsetEnpoint GoEndpointBackendManager.EndPoint) IReportStorageService {
	aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	err, ep := aepm.GetEndPoint()
	if err != nil {
		log.Println("Init Local TReportStorageService sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
		return &reportStorageService{
			hostEtcd:   defaultEnpoint.Host,
			portEtcd:   defaultEnpoint.Port,
			hostBigset: bigsetEnpoint.Host,
			portBigset: bigsetEnpoint.Port,
			sid:        defaultEnpoint.ServiceID,
		}
	}
	sv := &reportStorageService{
		hostEtcd:   ep.Host,
		portEtcd:   ep.Port,
		hostBigset: bigsetEnpoint.Host,
		portBigset: bigsetEnpoint.Port,
		sid:        ep.ServiceID,
	}
	go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	sv.epm = aepm
	log.Println("Init From Etcd TReportStorageService sid:", sv.sid, "host:", sv.hostEtcd, "port:", sv.portEtcd)
	log.Println("Init From Bigset TReportStorageService sid:", sv.sid, "host:", sv.hostEtcd, "port:", sv.portEtcd)

	return sv
}
