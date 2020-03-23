package MediaCloudEtcdClient

import (
	"context"
	"errors"

	"github.com/OpenStars/EtcdBackendService/MediaCloudEtcdClient/tmediacloudservice/thrift/gen-go/openstars/mcloud"

	"github.com/OpenStars/EtcdBackendService/MediaCloudEtcdClient/tmediacloudservice/transports"
)

type mediacloudclient struct {
	host string
	port string
}

func (m *mediacloudclient) GetMediaInfo(appId string, appKey string, mediaId string) (r *mcloud.TMCMediaInfoResult_, err error) {

	client := transports.GetMediaCloudServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}
	r, err = client.Client.(*mcloud.TMediaServiceClient).GetMediaInfo(context.Background(), appId, appKey, mediaId)
	if err != nil {
		return nil, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()
	return r, nil
}
