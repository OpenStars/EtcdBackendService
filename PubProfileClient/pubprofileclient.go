package PubProfileClient

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenStars/EtcdBackendService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	"github.com/OpenStars/EtcdBackendService/tpubprofileservice/transports"
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

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByUID(context.Background(), uid)
	if err != nil {
		return nil, errors.New("Backend service err:" + err.Error())
	}

	defer client.BackToPool()

	if resp != nil {
		fmt.Println(resp.ProfileData)
		return resp.ProfileData, nil
	}
	return nil, errors.New("Get data nil")
}

func (m *pubprofileclient) GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByPubkey(context.Background(), pubkey)
	if err != nil {
		return nil, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		return resp.ProfileData, nil
	}
	return nil, errors.New("Get data nil")
}

func (m *pubprofileclient) UpdateProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).UpdateProfileByPubkey(context.Background(), pubkey, profileUpdate)
	if err != nil {
		return false, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		return resp.Resp, nil
	}
	return false, nil
}

func (m *pubprofileclient) UpdateProfileByUID(uid int64, profileUpdate *pubprofile.ProfileData) (r bool, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).UpdateProfileByUID(context.Background(), uid, profileUpdate)
	if err != nil {
		return false, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		return resp.Resp, nil
	}
	return false, nil
}
