namespace go openstars.signservice

struct DataSignResult {
    1: string data;
    2: string pubkey;
    3: string sign;
    4: string errorMessage;
}

service SignDataService {
	DataSignResult SignData(1: string privateKey,2: string data);
}

