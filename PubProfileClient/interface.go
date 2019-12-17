package PubProfileClient

import "github.com/OpenStars/backendclients/go/tpubprofileservice/thrift/gen-go/openstars/pubprofile"

type PubProfileClientIf interface {
	GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error)
	GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error)
}
