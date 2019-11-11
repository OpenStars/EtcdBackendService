package SimpleSessionService

import "github.com/OpenStars/backendclients/go/simplesession/thrift/gen-go/simplesession"

type SimpleSessionClientIf interface {
	GetSession(sskey string) (*simplesession.TUserSessionInfo, error)
	CreateSession(uid int64) (sessionkey string, err error)
	RemoveSession(sskey string) error
}
