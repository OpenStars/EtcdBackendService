package MediaCloudEtcdClient

func NewMediaCloudEtcdClient(ahost, aport string) MediaCloudClientIf {

	return &mediacloudclient{
		host: ahost,
		port: aport,
	}
}
