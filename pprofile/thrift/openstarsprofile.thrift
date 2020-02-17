namespace cpp OpenStars.Platform.Profile
namespace go  OpenStars.Platform.Profile
namespace java openstars.platform.profile

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TSocialProfile{
    1: string sid,
    2: string name,
    3: string email,
}

struct TPlatformProfile{
    1: string username,
    2: string displayName
    3: map<string, bool> trustedEmails,
    4: map<string, bool> trustedMobiles,
    5: list<string> publicKeys, //for using with secp256k1
    6: map<string, string> ExtData,
    7: map<string, TSocialProfile> connectedSocial
}

typedef TPlatformProfile TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TPlatformProfile data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService extends TDataServiceR{
    TErrorCode putData(1: TKey key, 2: TPlatformProfile data)
    
}

service TPlatformProfileService extends TDataService{
    string setExtData(1: TKey uid, 2: string extKey, 3: string extValue),
        
    string getExtData(1: TKey uid, 2: string extKey),

    bool setTrustedEmail(1: TKey uid, 2: string email, 3: bool isTrusted),

    bool removeTrustedEmail(1: TKey uid, 2: string email),

    bool setTrustedMobile(1: TKey uid, 2: string email, 3: bool isTrusted),

    bool removeTrustedMobile(1: TKey uid, 2: string mobile),

    bool setSocialInfo(1: TKey uid, 2: string socialType, 3: TSocialProfile socialProfile),

    bool removeSocialInfo(1: TKey uid, 2: string socialType, 3: TSocialProfile socialProfile),
}


