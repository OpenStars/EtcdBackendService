package SimpleSessionService

import "github.com/OpenStars/EtcdBackendService/SimpleSessionService/simplesession/thrift/gen-go/simplesession"

type SimpleSessionClientIf interface {
	GetSession(sskey string) (*simplesession.TUserSessionInfo, error)
	CreateSession(uid int64) (sessionkey string, err error)
	RemoveSession(sskey string) error
}

type Client interface {
	GetSession(sskey string) (*simplesession.TUserSessionInfo, error)
	CreateSession(uid int64) (sessionkey string, err error)
	RemoveSession(sskey string) error

	// ================================= V2 ===========================================
	GetSessionV2(sskey string) (*simplesession.TUserSessionInfo, error)
	CreateSessionV2(uid int64, deviceInfo string, data string, expiredTime int64) (sessionkey string, err error)
	RemoveSessionV2(sskey string) (bool, error)
}
