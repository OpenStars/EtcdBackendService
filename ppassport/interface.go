package ppassport

import (
	"github.com/OpenStars/EtcdBackendService/ppassport/thrift/gen-go/OpenStars/Platform/Passport"
)

type PassportService interface {
	GetData(key int64) (*Passport.TPassportInfo, error)
	PutData(key int64, data *Passport.TPassportInfo) error
}
