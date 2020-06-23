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

struct TReportItem{
    1: i64 ID,
    2: string title,
    3: string content,
    4: i64 reportType,
    5: i64 uid,
    6: i64 timestamps,
    7: bool isDelete,
    8: map<string,string> mapExtend,
    9: i64 marketitemID,
}


typedef TReportItem TData


struct TListDataResult{
    1: TErrorCode errorCode,
    2: optional list<TReportItem> data
}

service TDataServiceR{
    TDataResult getData(1: i64 key), 
}
struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TReportItem data 
}


service TDataService{
    TDataResult getData(1:i64 key), 
    TErrorCode putData(1:i64 key, 2: TReportItem data)
    TErrorCode removeData(1:i64 key)
    TListDataResult getListData(1:list<i64> key)
}


service TReportItemService extends TDataService{
    
}


