namespace cpp OpenStars.Common.S2SKV
namespace go OpenStars.Common.S2SKV
namespace java OpenStars.Common.S2SKV

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TStringValue{
   1:TKey key,
   2:string value,
   3:i64 counter,
}

struct TPubkey2UidResult{
    1:TErrorCode errorCode,
    2:string pubkey,
    3:i64 uid,
}
struct TAddress2UidResult{
    1:TErrorCode errorCode,
    2:string address,
    3:i64 uid,
}
typedef TStringValue TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TStringValue data
}


service TDataServiceR{
}

service TDataService{
}

service TString2StringService extends TDataService{
    TPubkey2UidResult putPubkey2Uid(1: TKey pubkey),
    TErrorCode putAddress2Uid(1:TKey address,2:i64 uid),
    TAddress2UidResult getAddress2Uid(1:TKey address)
    TPubkey2UidResult getPubkey2Uid(1:TKey pubkey),
    TPubkey2UidResult getUid2Pubkey(1:i64 uid)
}


