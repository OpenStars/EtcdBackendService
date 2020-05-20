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

struct TCommentItem{
    1: TKey uid
    2: string username,
    3: string displayName
    4: map<string, bool> trustedEmails,
    5: map<string, bool> trustedMobiles,
    6: list<string> publicKeys, //for using with secp256k1

}

typedef TCommentItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TCommentItem data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TCommentItem data)
}

service TCommentStorageService extends TDataService{
    
}


