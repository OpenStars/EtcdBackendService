namespace cpp OpenStars.CallLog
namespace go OpenStars.calllog
namespace java OpenStars.CallLog

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TCallLog{
    1: string CallPhone,    
    2: string CallStatus,   
    3: string DeviceCall,  
    4: i64 ListenTime,   
    5: string PostOffice,   
    6: string PostmanId,    
    7: string PostmanName,  
    8: string ReceiverPhone,
    9: i64 RingTime,     
    10: string Type,         
    11: string CallType,    
    12: string Status,     
    13: string Source,      
    14: i64 Timestamp,    
    15: map<string, string> mapExt
}

typedef TCallLog TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TCallLog data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    map<TKey, TCallLog> getMultiData(1: list<TKey> keys), 
    TErrorCode putData(1: TKey key, 2: TCallLog data)
}

service TCallLogService extends TDataService{
    
}