namespace cpp OpenStars.EndUserVTP
namespace go openstars.enduservtp
namespace java OpenStars.EndUserVTP

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

enum TTypeUser{
    TType_Sender = 0,
    TType_Receiver = 1,
    TType_Other = 2,
}

struct TAddress {
    1: i64 addressID,
    2: string addressStr,
}

struct TEndUserVTP{
    1: TKey uid
    2: string phoneNumber,
    3: string displayName,
    4: TAddress address,
    7: string email,
    8: TTypeUser type,
    9: i64 evaluateUser,
    10: bool deleted,
    11: map<string, string> mapExtData,
    12: i64 createTime,
}

typedef TEndUserVTP TData

struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TEndUserVTP data
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    map<TKey, TEndUserVTP> getMultiData(1: list<TKey> keys),
    TErrorCode putData(1: TKey key, 2: TEndUserVTP data)
}

service TEndUserVTPService extends TDataService{
    
}


