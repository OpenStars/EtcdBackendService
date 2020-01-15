package PubProfileClient

import (
	"context"
	"errors"

	"github.com/OpenStars/backendclients/go/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	"github.com/OpenStars/backendclients/go/tpubprofileservice/transports"
)

type pubprofileclient struct {
	host string
	port string
}

func (m *pubprofileclient) GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	r, err = client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByUID(context.Background(), uid)
	if err != nil {
		return nil, errors.New("Backend service err:" + err.Error())
	}

	defer client.BackToPool()

	return r, nil
}

func (m *pubprofileclient) GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	r, err = client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByPubkey(context.Background(), pubkey)
	if err != nil {
		return nil, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	return r, nil
}
