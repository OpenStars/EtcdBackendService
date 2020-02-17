namespace cpp OpenStars.Common.TCommentStorageService
namespace go OpenStars.Common.TCommentStorageService
namespace java OpenStars.Common.TCommentStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}
struct ActionLink {
    1:optional string text,
    2:optional string href,
}

struct MediaItem {
    1:string name,
    2:i64 mediaType, // 1 = image ; 2 = video; 3 = gif; 
    3:string url,
}

struct TCommentItem{
    1: TKey idcomment,
    2: i64 uid,
    3: string pubkey,
    4: i64 idpost,
    5: string content,
    6: optional list<ActionLink> actionlinks,
    7: optional MediaItem mediaitem,
    8: optional list<string> pubkeyTags,
    9: optional list<string> uidTags,
    10: i64 timestamps,
    11: optional i64 parentcommentid,
    12: map<string,string> mapExtend,
}

typedef TCommentItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TCommentItem data
    
}
struct TListDataResult {
    1: TErrorCode errorCode,
    2: list<TCommentItem> listDatas;
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TCommentItem data)
    TListDataResult getListData(1:list<TKey> listKey)
    TErrorCode removeData(1:TKey key)
}

service TCommentStorageService extends TDataService{
    
}


