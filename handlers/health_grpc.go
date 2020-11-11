package handlers

import (
	"context"

	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"

	hpb "github.com/wksw/protos-repo/health"
)

func init() {
	chassis.RegisterSchema("grpc", &Health{}, server.WithRPCServiceDesc(&hpb.Health_serviceDesc))
}

type Health struct{}

var _ hpb.HealthServer = &Health{}

// 健康检查
func (hlt *Health) Check(ctx context.Context, in *hpb.HealthCheckRequest) (*hpb.HealthCheckResponse, error) {
	return &hpb.HealthCheckResponse{
		Status: hpb.HealthCheckResponse_SERVING,
	}, nil
}
