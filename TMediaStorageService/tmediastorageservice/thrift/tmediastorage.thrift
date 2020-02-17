namespace cpp OpenStars.Common.TMediaStorageService
namespace go OpenStars.Common.TMediaStorageService
namespace java OpenStars.Common.TMediaStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TMediaItem {
    1:string name,
    2:i64 mediaType, // 1 = image ; 2 = video; 3 = gif; 
    3:string url,
    4:i64 idmedia // == idpost
    5:i64 idpost,
    6:i64 timestamps,
    7:string extend,
    8:map<string,string> mapExtend,
}

typedef TMediaItem TData


struct TDataResult{ 
    1: TErrorCode errorCode,
    2: optional TMediaItem data
    
}

struct TListDataResult {
    1: TErrorCode errorCode,
    2: list<TMediaItem> listDatas;
}

service TDataServiceR{
    TDataResult getData(1: TKey key),
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TMediaItem data)
    TErrorCode removeData(1:TKey key),
    TListDataResult getListData(1:list<TKey> listkey),
}

service TMediaStorageService extends TDataService{
    
}