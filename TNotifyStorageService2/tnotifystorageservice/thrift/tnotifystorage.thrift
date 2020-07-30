namespace cpp OpenStars.Common.TNotifyStorageService
namespace go OpenStars.Common.TNotifyStorageService
namespace java OpenStars.Common.TNotifyStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TNotifyItem{
    1: i64 Id, // notifyitem id
    2: i64 subjectId // chu the phat ra hanh dong
    3: string actionType, // hanh dong
    4: string objectId, // doi tuong chiu tac dong cua chu the
    5: map<string,string> messLang, // render message in language
    6: bool seen,
    7: i64 timestamps,
    8: map<string,string> mapExtend,
    9: i64 uid, // nguoi nhan notification
}

typedef TNotifyItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TNotifyItem data  
}

struct TListDataResult{
    1:TErrorCode errorCode,
    2:optional list<TNotifyItem> datass,
}

service TDataServiceR{
    TDataResult getData(1:i64 key), 
}

service TDataService{
    TDataResult getData(1:i64 Id), 
    TErrorCode putData(1:i64 Id, 2:TNotifyItem data),
    TListDataResult getListData(1:list<i64> lskeys)
    bool removeData(1:i64 id)
}

service TNotifyStorageService extends TDataService{
    
}



