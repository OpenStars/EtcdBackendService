package MediaCloudEtcdClient

func NewPubProfileClient(ahost, aport string) MediaCloudClientIf {

	return &mediacloudclient{
		host: ahost,
		port: aport,
	}
}
