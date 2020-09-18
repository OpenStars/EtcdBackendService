namespace cpp OpenStars.Order
namespace go OpenStars.orderservice
namespace java OpenStars.Order

typedef string TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TOrder{
    1: TKey order_ladingcode,
    2: i64 order_createdate,
    3: string order_sendname,
    4: string order_sendaddress,
    5: string order_sendtel,
    6: string order_sendprovince,
    7: string order_description,
    8: string order_note,
    9: i64 ngay_giao,
    10: string buucucgo,
    11: i64 buu_ta,
    12: i64 trang_thai,
    13: i64 id_nhan,
    14: string mota_sp,
    15: string ma_dv_viettel,
    16: i64 trong_luong,
    17: i64 tien_hang,
    18: i64 tien_thu_ho,
    19: string order_reference,
    20: string dv_cong_them,
    21: i64 thoigian,
    22: string ma_buucuc,
    23: i64 dieu_hanh,
    24: string order_sendward,
    25: string order_senddistrict,
}

typedef TOrder TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TOrder data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key),
    map<TKey, TOrder> getMultiData(1: list<TKey> keys),
    TErrorCode putData(1: TKey key, 2: TOrder data)
}

service TOrderService extends TDataService{
    
}


