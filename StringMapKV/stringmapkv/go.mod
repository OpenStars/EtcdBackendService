module github.com/OpenStars/backendclients/go/stringmapkv

go 1.13

require git.apache.org/thrift.git/lib/go/thrift v0.0.0

replace git.apache.org/thrift.git/lib/go/thrift v0.0.0 => /home/lehaisonmath6/go/src/git.apache.org/thrift.git/lib/go/thrift

require (
	github.com/OpenStars/thriftpool v0.0.0
	github.com/OpenStars/thriftpoolv2 v0.0.0-20191121022147-482e96ec8e92
	github.com/apache/thrift v0.13.0
)

replace github.com/OpenStars/thriftpool v0.0.0 => /home/lehaisonmath6/go/src/github.com/OpenStars/thriftpool

// require github.com/OpenStars/backendclients/go/stringmapkv/thrift/gen-go/OpenStars/Common/StringMapKV v0.0.0

// replace github.com/OpenStars/backendclients/go/stringmapkv/thrift/gen-go/OpenStars/Common/StringMapKV v0.0.0 => /home/lehaisonmath6/go/src/github.com/OpenStars/backendclients/go/stringmapkv/thrift/gen-go/OpenStars/Common/StringMapKV
