namespace cpp OpenStars.Order
namespace go OpenStars.orderservice
namespace java OpenStars.Order

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TOrder{
    1: TKey orderGenID,
    2: i64 order_id,
    3: string order_number,
    4: string order_reference,
    5: i64 groupaddress_id,
    6: i32 partner,
    7: i64 delivery_date,
    8: string sender_fullname,
    9: string sender_address,
    10: string sender_phone,
    11: i32 sender_ward,
    12: i32 sender_district,
    13: i32 sender_province,
    14: string receiver_address,
    15: string receiver_phone,
    16: i32 receiver_ward,
    17: i32 receiver_district,
    18: i32 receiver_province,
    19: string product_name,
    20: i32 product_price,
    21: i32 product_weight,
    22: string product_type,
    23: i32 order_payment,
    24: string order_service,
    25: string order_service_add,
    26: i32 order_status,
    27: i32 order_post_id,
    28: i64 order_systemdate,
    29: i32 order_employer,
    30: string order_note,
    31: i64 money_collection,
    32: i64 money_totalfee,
    33: i64 money_feecod,
    34: i64 money_feevas,
    35: i64 money_feeinsurrance,
    39: i64 money_fee,
    40: i64 money_feeother,
    41: i64 money_totalvat,
    42: i64 money_total,
    43: i64 order_type,
    44: string post_code,
    45: string service_name,
    46: string province_code,
    47: string district_name,
    48: string district_code,
    49: string wards_code,
    50: i32 is_pending,
    51: i32 order_action_505
    52: i32 fee_collected,
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
    TErrorCode putData(1: TKey key, 2: TOrder data)
}

service TOrderService extends TDataService{
    
}


