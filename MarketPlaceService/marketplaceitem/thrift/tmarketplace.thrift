namespace cpp OpenStars.Platform.MarketPlace
namespace go OpenStars.Platform.MarketPlace
namespace java OpenStars.Platform.MarketPlace

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TMediaItem{
    1:string imgurl,
    2:string name,
    3:i64   mediaType,
    4:map<string,string> extend,
}

struct TLocation {
    1:double latitude,
    2:double longitude,
    3:map<string,string> extend,
}
struct TMarketPlaceItem{
    1: i64 ID,
    2: string title,
    3: i64 price,
    4: list<TMediaItem> listMediaItems,
    5: i64 category,
    6: map<string,string> subfeatures,
    7: string descriptions,
    8: i64 uid,
    9: i64 count,
    10: bool isdelivery,
    11: list<string> tags,
    12: i64 timestamps,
    13: TLocation location,
    14: bool isDelete,
    15: map<string,string> mapExtend,
}

typedef TMarketPlaceItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TMarketPlaceItem data
    
}

struct TListDataResult{
    1: TErrorCode errorCode,
    2: optional list<TMarketPlaceItem> data
}

service TDataServiceR{
    TDataResult getData(1: i64 key), 
}

service TDataService{
    TDataResult getData(1:i64 key), 
    TErrorCode putData(1:i64 key, 2: TMarketPlaceItem data)
    TErrorCode removeData(1:i64 key)
    TListDataResult getListData(1:list<i64> key)
}

service TMarketPlaceService extends TDataService{
    
}


