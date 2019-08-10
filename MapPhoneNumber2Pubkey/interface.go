package MapPhoneNumber2Pubkey

type MappingPhoneNumber2PubkeyServiceIf interface {
	PutData(pubkey string, phonenumber string) error
	GetPhoneNumberByPubkey(pubkey string) (string, error)
	GetPubkeyByPhoneNumber(phonenumber string) (string, error)
}
