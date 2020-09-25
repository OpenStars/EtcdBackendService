namespace cpp OpenStars.VTPComment
namespace go OpenStars.VTPComment
namespace java OpenStars.VTPComment

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TVTPComment{
    1: TKey id,
    2: string text,
    3: string uid_comment, 
    4: i64 time,
    5: map<string, bool> mapExt,
    6: string ma_buugui,
}

typedef TVTPComment TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TVTPComment data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TVTPComment data)
    map<TKey, TVTPComment> getMultiData(1: list<TKey> keys)
}

service TVTPCommentService extends TDataService{
    
}


