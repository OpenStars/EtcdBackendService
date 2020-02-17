namespace cpp OpenStars.Common.TReportStorageService
namespace go OpenStars.Common.TReportStorageService
namespace java OpenStars.Common.TReportStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct MediaItem {
    1:string name,
    2:i64 mediaType, // 1 = image ; 2 = video; 3 = gif;
    3:string url,
}

struct TReportItem{
    1: i64 reportId
    2: i64 uId,
    3: i64 targetId,
    4: optional i64 postId,
    5: optional i64 commentId,
    6: optional string contentObj,
    7: optional list<MediaItem> listMediaObj,
    8: optional list<string> actionLink,
    9: optional string contentReport,
    10: i64 timestamp,
    11: optional i64 timestampObj,
    12: optional string locationId
    13: optional string locationName
    14: i8 action,
}

typedef TReportItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TReportItem data  
}

struct TListDataResult{
    1:TErrorCode errorCode,
    2:optional list<TReportItem> data,
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}


service TDataService{
    TDataResult getReport(1:i64 reportId), 
    TErrorCode putReport(1:i64 reportId, 2: TReportItem data)
    bool RemoveReport(1:i64 reportId)
    TListDataResult getListReports(1:list<i64> lsReportIds)

}

service TReportStorageService extends TDataService{
    
}


