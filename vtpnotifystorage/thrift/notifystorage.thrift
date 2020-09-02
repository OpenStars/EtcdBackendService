namespace cpp OpenStars.NotifyStorage
namespace go OpenStars.notifystorage
namespace java OpenStars.NotifyStorage

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TNotifyStorage{
    1: TKey id,
    2: i64 senderID,
    3: i64 receiverID,
    4: string action,
    5: bool seen,
    6: map<string, string> mapExt,
    7: bool deleted,
    8: string app,
    9: string icon,
    10: i64 timestamp,
    11: string title, 
    12: string content,
}

typedef TNotifyStorage TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TNotifyStorage data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TNotifyStorage data)
    map<TKey, TNotifyStorage> getMultiData(1: list<TKey> keys)
}

service TNotifyStorageService extends TDataService{
    
}


