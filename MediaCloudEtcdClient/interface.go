package MediaCloudEtcdClient

import "github.com/OpenStars/EtcdBackendService/MediaCloudEtcdClient/tmediacloudservice/thrift/gen-go/openstars/mcloud"

type MediaCloudClientIf interface {
	GetMediaInfo(appId string, appKey string, mediaId string) (r *mcloud.TMCMediaInfoResult_, err error)
}
