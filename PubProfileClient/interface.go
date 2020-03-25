package PubProfileClient

import "github.com/OpenStars/EtcdBackendService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"

type PubProfileClientIf interface {
	GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error)
	GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error)
	UpdateProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error)
	UpdateProfileByUID(uid int64, profileUpdate *pubprofile.ProfileData) (r bool, err error)
}
