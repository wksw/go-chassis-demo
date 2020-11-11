package handlers

import (
	"context"

	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"

	pb "github.com/wksw/protos-repo/hello"
)

func init() {
	chassis.RegisterSchema("grpc", &HelloGrpc{}, server.WithRPCServiceDesc(&pb.Hello_serviceDesc))
}

type HelloGrpc struct{}

var _ pb.HelloServer = &HelloGrpc{}

func (h *HelloGrpc) Hello(c context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{
		Resp: in.Req,
	}, nil
}
