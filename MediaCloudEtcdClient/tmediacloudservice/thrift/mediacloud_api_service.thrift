include "mediacloud_shared.thrift"

namespace cpp mcloud
namespace java com.bt.media.cloud.thrift
namespace go openstars.mcloud

service TMediaService {

    void ping();

	mediacloud_shared.TMCProcessResult process(1:string appId, 2:string appKey, 3: mediacloud_shared.TMCProcessOption option);
	
	//get info
	mediacloud_shared.TMCMediaInfoResult getMediaInfo(1:string appId, 2:string appKey, 3:string mediaId);
}
