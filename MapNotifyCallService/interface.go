package MapNotifyCallService

type MapNotifyCallServiceIf interface {
	PutData(pubkey string, token string) error
	GetTokenByPubkey(pubkey string) (string, error)
	GetPubkeyByToken(token string) (string, error)
}
