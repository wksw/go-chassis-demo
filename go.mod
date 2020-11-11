module github.com/wksw/go-chassis-demo

go 1.13

require (
	github.com/go-chassis/go-archaius v1.3.6-0.20200917065837-57a2bca2b7ff
	github.com/go-chassis/go-chassis-extension/protocol/grpc v0.0.0-20200803124338-47eb21ed82a7
	github.com/go-chassis/go-chassis/v2 v2.0.2
	github.com/go-chassis/openlog v1.1.2
	github.com/go-chassis/seclog v1.3.0 // indirect
	// github.com/go-chassis/seclog/third_party/forked/cloudfoundry/lager v0.0.0 // indirect
	github.com/wksw/protos-repo v0.0.0-00010101000000-000000000000
)

// github.com/go-chassis/go-archaius => /go/src/github.com/wksw/go-archaius
// github.com/go-chassis/go-chassis-extension/protocol/grpc => /go/src/github.com/wksw/go-chassis-extension/protocol/grpc
// github.com/go-chassis/go-chassis/v2 => /go/src/github.com/wksw/go-chassis
// github.com/go-chassis/openlog => /go/src/github.com/wksw/openlog
// github.com/go-chassis/seclog => /go/src/github.com/wksw/seclog
// github.com/go-chassis/seclog/third_party/forked/cloudfoundry/lager => /go/src/github.com/wksw/seclog/third_party/forked/cloudfoundry/lager
replace github.com/wksw/protos-repo => ./thirdparty/protos-repo
