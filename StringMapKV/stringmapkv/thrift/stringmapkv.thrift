namespace cpp OpenStars.Common.StringMapKV
namespace go OpenStars.Common.StringMapKV
namespace java OpenStars.Common.StringMapKV

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TStringValue{
    1:string value
}

typedef TStringValue TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TStringValue data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TStringValue data)
}

service StringMapKVService extends TDataService{
    
}


