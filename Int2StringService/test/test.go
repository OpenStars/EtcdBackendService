package main

import (
	"github.com/OpenStars/EtcdBackendSerivce/StringBigsetSerivce"

	"github.com/OpenStars/EtcdBackendSerivce/Int2StringService"
	"github.com/OpenStars/EtcdBackendSerivce/String2Int64Service"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type Server struct {
	int2string   Int2StringService.Int2StringServiceIf
	string2int   String2Int64Service.String2Int64ServiceIf
	stringbigset StringBigsetSerivce.StringBigsetServiceIf
}

func (s *Server) Run() {
	for {

	}
}

func TestStringBigset() {
	sid := "/trustkeys/tkverifyprofile/stringbigset"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := StringBigsetSerivce.NewStringBigsetServiceModel(sid, etcd, defaultEp)
	sv := &Server{
		stringbigset: ai2s,
	}
	sv.Run()
}

func TestString2Int() {
	sid := "/openstars/services/string2int"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := String2Int64Service.NewString2Int64Service(sid, etcd, defaultEp)
	sv := &Server{
		string2int: ai2s,
	}
	sv.Run()
}

func TestInt2String() {
	sid := "/openstars/services/int2string"
	etcd := []string{"127.0.0.1:2379"}
	defaultEp := GoEndpointBackendManager.EndPoint{
		ServiceID: sid,
		Host:      "127.0.0.1",
		Port:      "8883",
	}
	ai2s := Int2StringService.NewInt2StringService(sid, etcd, defaultEp)
	sv := &Server{
		int2string: ai2s,
	}
	sv.Run()
}

func main() {
	// TestInt2String()
	// TestString2Int()

	TestStringBigset()
}
