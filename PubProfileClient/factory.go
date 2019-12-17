package PubProfileClient

func NewPubProfileClient(ahost, aport string) PubProfileClientIf {

	return &pubprofileclient{
		host: ahost,
		port: aport,
	}
}
