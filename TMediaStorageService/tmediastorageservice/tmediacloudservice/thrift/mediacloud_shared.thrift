namespace cpp mcloud
namespace java com.bt.media.cloud.thrift
namespace php mcloud
namespace go openstars.mcloud

enum TMCVideoFormat {
	MP4_H264 	= 2,
	//AVI 		= 3,
	//MOV 		= 4,
	WEBM_VP8 	= 5,
	//WMV 		= 6,
	//TS 		= 7,
	HLS 		= 8,
	WEBM_VP9 	= 9,
	MP4_H265 	= 10,
	DASH_VP9	= 11,
	//_3GP 		= 12,
	HLS_V2		= 13,
	HLS_WRAP_MP4	= 14,
	DASH_H265	= 15,
	HLS_H265	= 16,
	HLS_WRAP_MP4_SC	= 17,
}

enum TMCVideoQuality {
	QUALITY_UNKNOWN = 0,
	QUALITY_240P 	= 1,
	QUALITY_360P 	= 2,
	QUALITY_480P 	= 3,
	QUALITY_720P 	= 4,
	QUALITY_1080P 	= 5,
	QUALITY_1440P 	= 6,
	QUALITY_2160P 	= 7,
	QUALITY_144P	= 8,
}

enum TMCImageQuality {
	QUALITY_UNKNOWN = 0,
	QUALITY_240P 	= 1,
	QUALITY_360P 	= 2,
	QUALITY_480P 	= 3,
	QUALITY_720P 	= 4,
	QUALITY_144P	= 8,
}

enum TMCGifQuality {
	QUALITY_UNKNOWN = 0,
	QUALITY_240P 	= 1,
	QUALITY_360P 	= 2,
	QUALITY_480P 	= 3,
	QUALITY_720P 	= 4,
	QUALITY_144P	= 8,
}

enum TMCAudioQuality {
	QUALITY_UNKNOWN 	= 0,
	QUALITY_32K 		= 1,
	QUALITY_64K 		= 2,
	QUALITY_96K 		= 3,
	QUALITY_128K 		= 4,
	QUALITY_256K 		= 5,
	QUALITY_320K 		= 6,
	QUALITY_500K 		= 7,
	QUALITY_LOSSLESS	= 8,
}

enum TMCMediaStatus {
	QUEUING 		= 0,
	PROCESSING 		= 1,
	DONE 			= 2,
	DONE_WITH_ERROR 	= 3,
	DOWNLOADING 		= 4,
	INPUT_FILE_INVALID 	= 100,
	ID_ZEN_ERROR 		= 101,
	PROCESS_ERROR 		= 102,
	DATABASE_ERROR 		= 103,
	UNSUPPORTED_FORMAT 	= 104,
	UNSUPPORTED_QUALITY 	= 105,
	QLT_HIGHER_THAN_ORIGIN 	= 106,
	DOWNLOAD_INTERNAL_FAIL 	= 107,
	DOWNLOAD_EXTERNAL_FAIL 	= 108,
	NETWORK_ERROR 		= 109,
	DELETED 		= 110,
    INPUT_QLT_LOW		= 111,
}

enum TMCSourceType {
	FTP 	= 0,
	HTTP 	= 1,
	LOCAL 	= 2,
}

struct TMCSourceInfo {
	1:optional i32 sourceType, //ref TMCSourceType
	2:optional string path,
	3:optional bool isExternalSource,
    4:optional i64 size,
	5:optional i32 duration,
	6:optional bool isAnimationVideo,
}

struct TMCWaterMarkInfo {
	1:optional TMCSourceInfo sourceInfo,
	2:optional i32 marginLeft,
	3:optional i32 marginTop,
}

struct TMCOutputInfo {
	1:optional map<i32, set<i32>> formats, // map<format, set<quality>>
}

struct TMCStoryBoardId {
	// status
	1:optional i32 status, // ref TMCMediaStatus
	2:optional i32 remainTime,
	3:optional i32 totalTimeProc,
	// info
	10:optional i16 period,
	11:optional i16 imageRes,
	12:optional byte boardSize,
	13:optional list<i64> imageIds,
}

struct TMCStoryBoardUrl {
	// status
	1:optional i32 status, // ref TMCMediaStatus
	2:optional i32 remainTime,
	3:optional i32 totalTimeProc,
	// info
	10:optional i16 period,
	11:optional i16 imageRes,
	12:optional byte boardSize,
	13:optional list<string> imageUrls,
}

struct TMCPosterSource {
	1:optional bool autoGen,
	2:optional set<i32> quality, // ref TMCImageQuality
	3:optional list<TMCSourceInfo> sources, // if manual
}

struct TMCGifSource {
	1:optional bool autoGen,
	2:optional set<i32> quality, // ref TMCGifQuality
	3:optional bool fromImage,
	4:optional list<TMCSourceInfo> sources, // if manual
}

struct TMCProcessOption {
	// input
	1:optional string mediaName,
	2:optional string appId,
	3:optional TMCSourceInfo mediaSource,
	//
	10:optional bool addWatermark,
	11:optional TMCWaterMarkInfo watermark,
	// storyBoard
	20:optional bool genStoryboard,
	// poster
	30:optional bool genPoster,
	31:optional TMCPosterSource posterSource,
	// posterGif
	40:optional bool genGif,
	41:optional TMCGifSource posterGifSource,
	// output
	50:optional TMCOutputInfo output,
}

struct TMCProcessResult {
	1:required i32 error, //ref zcommon.ECode
	2:optional string mediaId,
}

struct TMCFileStatusId {
	1:optional i32 status, // ref TMCMediaStatus
	2:optional i64 fileId,
	3:optional i32 remainTime,
	4:optional i32 totalTimeProc,
    5:optional i64 createTime,
	6:optional i32 storageId,
	7:optional binary extraData,
	8:optional i64 fileSize,
	9:optional i16 videoWidth,
	10:optional i16 videoHeight
}

struct TMCFileStatusUrl {
	1:optional i32 status, // ref TMCMediaStatus
	2:optional string fileUrl,
	3:optional i32 remainTime,
	4:optional i32 totalTimeProc,
	5:optional i64 fileSize,
	6:optional i32 videoWidth,
	7:optional i32 videoHeight
}

struct TMCMediaProperties {
	1:optional i32 duration,
}

typedef map<i32, TMCFileStatusUrl> TMCQualityInfoMap

struct TMCMediaInfo {
	1:optional i32 status, //ref TMCMediaStatus
	2:optional map<i32, TMCQualityInfoMap> mapFormatQualityInfo, // ref <TMCVideoFormat - <TMCVideoQuality - TMCFileStatusUrl> >
	//
	6:optional string mediaName,
	7:optional i64 createTime,
	8:optional i32 totalTimeProc,
	9:optional i32 downloadProgress,
	10:optional i64 originFileSize,
	11:optional TMCSourceInfo sourceInfo,
	12:optional TMCSourceInfo externalSourceInfo,
	13:optional i32 duration,
	14:optional TMCMediaProperties properties,
	//
	15:optional string appId,
	16:optional string tempOriginUrl,
	//
	30:optional TMCStoryBoardUrl storyBoardUrls,
	31:optional list<map<i32, TMCFileStatusUrl>> posterUrls,
	32:optional list<map<i32, TMCFileStatusUrl>> posterGifUrls,
}

struct TMCMediaMeta {
	1:required i64 mediaId,
	2:optional i32 status, //ref TMCMediaStatus
	3:optional map<i32, map<i32, TMCFileStatusId>> mapFormatInfo, // ref <TMCVideoFormat - <TMCVideoQuality - TMCFileStatusId> >
	4:optional list<i64> listPosterId, //poster id (bo sau khi off mw cu)
	// 5:optional list<i64> listStoryboardId,
	6:optional string mediaName,
	7:optional i64 createTime,
	8:optional i32 totalTimeProc,
	9:optional i32 downloadProgress,    
	10:optional i64 originFileSize,
	11:optional TMCSourceInfo sourceInfo,
	12:optional TMCSourceInfo externalSourceInfo,
	13:optional i32 duration,
	// 14:optional list<i64> listPosterGifId,
	15:optional i32 appId,
	16:optional i32 storageId,
	17:optional string tempOriginUrl,
	18:optional TMCMediaProperties properties,
	// extend media
	30:optional TMCStoryBoardId storyBoardIds,
	31:optional list<map<i32, TMCFileStatusId>> posterIds, // ref list(map<i32(quality in number),ImageID>)
	32:optional list<map<i32, TMCFileStatusId>> posterGifIds, // ref list(map<i32(quality in number),GifID>)
}

struct TMCMediaInfoResult {
	1:required i32 error; //ref ECode
	2:optional TMCMediaInfo value;
}

struct TMCMediaMetaResult {
	1:required i32 error,
	2:optional TMCMediaMeta data,
}

// for worker convert
struct TMCMediaProcessNotify {
	1:required i64 mediaId,
	2:optional TMCSourceInfo mediaSource,
	// xoa sau khi deploy code moi
	  3:optional TMCWaterMarkInfo watermarkInfo,
	  4:optional TMCOutputInfo outputInfo,
	  5:optional bool zenPoster,
	  6:optional bool zenStoryboard,
	  8:optional bool getDuration, // khong can check, default update duration
	// convert
	10:optional bool convert,
	11:optional TMCWaterMarkInfo watermark,
	12:optional TMCOutputInfo output,
	// storyBoard
	20:optional bool genStoryboard,
	// poster
	30:optional bool genPoster,
	31:optional TMCPosterSource posterSource,
	// poster gif
	40:optional bool genGif,
	41:optional TMCGifSource posterGifSource,

    50: optional i32 appId;
	
}

// for worker update status
struct TMCProcessStatusNotify {
	1:required i64 mediaId,
	2:optional i32 status, // ref TMCMediaStatus
	// xoa sau khi deploy code moi
	  3:optional map<i32, map<i32, TMCFileStatusId>> mapFormatProgress_obs, // map<format, map<quality, TMCFileStatusId>>
	  4:optional list<i64> listPosterId_obs, //poster id
	  5:optional list<i64> listStoryboardId_obs,
	  6:optional i32 downloadProgress_obs,
	  7:optional i32 totalTimeProc_obs,
	  8:optional i32 duration_obs,
	// convert
	10:optional bool updateConvertProgress,
	11:optional map<i32, map<i32, TMCFileStatusId>> mapFormatProgress, // map<format, map<quality, TMCFileStatusId>>,
	12:optional i32 duration,
	13:optional i64 originFileSize,
	// download
	20:optional bool updateDownloadProgress,
	21:optional i32 downloadProgress,
	// storyBoard
	30:optional bool updateStoryBoardProgress,
	31:optional TMCStoryBoardId storyBoardIds,
	// poster
	40:optional bool updatePosterProgress,
	41:optional list<map<i32, TMCFileStatusId>> posterIds,
	// posterGif
	50:optional bool updatePosterGifProgress,
	51:optional list<map<i32, TMCFileStatusId>> posterGifIds,
	
}

struct TMCDownloadExternalNotify {
	1:optional i64 mediaId,
	2:optional TMCProcessOption processOption,
}

struct TMCMediaProfile {
	1:optional string id,
	2:optional i32 videoBitrateKbps,
	3:optional i32 fps,
	4:optional i32 resolution,
	5:optional i32 audioBitrateKbps,
}

struct TMCFileSizeTranscode {
	1:optional i32 fileSize,
	2:optional i32 farmId,
}

// for admin
struct TMCAppMedia {
	1: required i32 appId,
	2: required string appName,
	3: required string appKey,
	4: optional i32 appType,
	5: optional i32 storageId,
	6: optional string domain,
	7: optional i32 farmProcessId,
	//8: optional i32 minExpireTime,
	//9: optional i32 maxExpireTime,
	10:optional map<i32, map<i32, i32>> mapFormatQualityStorage,
	//11 has been used for some thing else, do not define 
	21:optional map<string, TMCMediaProfile> mapProfile,
	22:optional list<TMCFileSizeTranscode> farmProcessBySize,
	23:optional i32 secureLinkExpireBf,
	24:optional i32 secureLinkExpireHls,
	25:optional i32 storageForMaxQuality,
	26:optional i32 storageForMinQuality,
	27:optional i32 videoKeyframeInterval,
	28:optional bool useOriginDomain,
	29:optional bool useHttps,
	
	30: optional map<i32, set<i32>> appPermission, // map from format to set of qualities
	//31: optional map<i32, map<i32, i32>> appStorage, // map from format => map<quality, storage>
	//32: optional i32 defaultBitrate, // unit Mbps
	//33: optional map<i32, i32> bitrateConfig, // map from Quality Id to bitrate
	34: optional bool isInternalApp,
	35:optional i32 minInputQuality,
	36:optional i32 minInputAudioQuality,
	
	37:optional i32 videoPreset,
	
}

//storage

struct THostInfo {
	1:optional string address,
	2:optional i32 port,
}

struct TServiceInfo {
	// read
	1:optional list<THostInfo> readHosts,
	2:optional i32 readScaleMode,
	3:optional string readSource,
	4:optional string readAuth,
	5:optional i32 readTimeout,
	// write
	20:optional list<THostInfo> writeHosts,
	21:optional i32 writeScaleMode,
	22:optional string writeSource,
	23:optional string writeAuth,
	24:optional i32 writeTimeout,
	//config
	30:optional THostInfo cfgHost,
}

struct TChunkInfo {
	1:optional i32 chunkId,
	2:optional TServiceInfo serviceInfo,

	20:optional string name,
}

struct TMCHLSInfo {
	//1:optional zstorage_meta.TServiceInfo metaInfo,
	//2:optional zstorage_meta.TServiceInfo playlistInfo,
	//3:optional zstorage_meta.TServiceInfo subPlaylistInfo,
	4:optional list<i32> allChunks,
	5:optional list<i32> writeChunks,
	//6:optional string masterDomain,
	//7:optional string levelDomain,
	//8:optional string chunkDomain,
	9:optional i64 playlistKeyNoise,
	10:optional i64 chunkKeyNoise,
	//11:optional string levelZdnAppName,
	//12:optional string chunkZdnAppName,
	13:optional i16 writeMode,
	
	14:optional TServiceInfo playlistV2Info,
	15:optional string masterDomainV2,
	16:optional string levelDomainV2,
	17:optional string chunkDomainV2,
	18:optional string chunkDomainWrapper,

	//20:optional i32 id,
	//21:optional string name,
	
	22:optional string masterOrgDomain,
	23:optional string levelOrgDomain,
}

struct TMCBFInfo {
	1:optional TServiceInfo metaInfo,
	2:optional list<i32> allChunks,
	3:optional list<i32> writeChunks,
	4:optional string domain,
	5:optional i64 keyNoise,
	6:optional i16 writeMode,
	7:optional string orgDomain,
}

struct TMCImageInfo {
	1:optional TServiceInfo metaInfo,
	2:optional list<i32> allChunks,
	3:optional list<i32> writeChunks,
	4:optional string domain,
	5:optional i64 keyNoise,
	6:optional i16 writeMode,
	7:optional string orgDomain,
}

struct TMCDASHInfo {
	//1:optional zstorage_meta.TServiceInfo playlistMetaInfo,
	2:optional TServiceInfo fileMetaInfo,
	3:optional list<i32> allChunks,
	4:optional list<i32> writeChunks,
	5:optional string playlistDomain,
	6:optional string fileDomain,
	7:optional i64 playlistKeyNoise,
	8:optional i64 fileKeyNoise,
	9:optional i16 writeMode,
	10:optional string playlistOrgDomain
}

struct TMCStorageInfo {
	1:optional i32 storageId,
	2:optional string name,
	3:optional TMCHLSInfo hlsInfo,
	4:optional TMCBFInfo bigFileInfo,
	5:optional TMCImageInfo imageInfo,
	6:optional TMCDASHInfo dashInfo,
	7:optional i32 lowBoundExpireTime,
	8:optional i32 highBoundExpireTime,
}

// farm process info
struct ZMCFarmProcessInfo {
	1:optional i32 farmId,
	2:optional string name,
	3:optional TServiceInfo eventBusInfo,
	//video
	10:optional string h264Event,
	11:optional string h265Event,
	12:optional string vp8Event,
	13:optional string vp9Event,
	14:optional string hlsEvent,
	15:optional string dashEvent,
	//audio
	20:optional string audioEvent,
	
}