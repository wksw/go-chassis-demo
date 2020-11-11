package handlers

import (
	"net/http"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/common"
	"github.com/go-chassis/go-chassis/v2/core/server"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"

	hpb "github.com/wksw/protos-repo/health"
	pb "github.com/wksw/protos-repo/hello"
)

func init() {
	chassis.RegisterSchema("rest", &Hello{}, server.WithSchemaID(SERVICE_NAME))
}

type Hello struct{}

type Say struct {
	Say string `json:"say" form:"say"`
}

// 无参数解析
func (h *Hello) Say(b *rf.Context) {
	response(b,
		&Say{
			Say: "world",
		},
		nil)
}

// query参数
func (h *Hello) QuerySay(b *rf.Context) {
	var req Say
	if err := b.ReadQueryEntity(&req); err != nil {
		openlog.Errorf("query fail[%s]", err.Error())
		response(b, &Say{}, err)
		return
	}
	response(b, &req, nil)
}

// path参数
func (h *Hello) PathSay(b *rf.Context) {
	var req Say
	if err := b.ReadQueryEntity(&req); err != nil {
		response(b, &Say{}, err)
		return
	}
	response(b, &req, nil)
}

// body参数
func (h *Hello) BodySay(b *rf.Context) {
	var req Say
	if err := b.ReadEntity(&req); err != nil {
		response(b, &Say{}, err)
		return
	}
	response(b, &req, nil)
}

// query、path、body参数
func (h *Hello) BodyQueryPathSay(b *rf.Context) {
	var req Say
	if err := b.ReadEntity(&req); err != nil {
		response(b, &Say{}, err)
		return
	}
	response(b, &req, nil)
}

func (h *Hello) HelloGrpc(b *rf.Context) {
	var req pb.HelloReq
	if err := b.ReadQueryEntity(&req); err != nil {
		response(b, &pb.HelloResp{}, err)
		return
	}
	var headers = make(map[string]string)
	ctx := common.NewContext(headers)
	client := pb.NewHelloClient(ctx, archaius.GetString("servicecomb.service.name", SERVICE_NAME))
	resp, err := client.Hello(&req)
	response(b, resp, err)
}

func (h *Hello) Health(b *rf.Context) {
	var req hpb.HealthCheckRequest
	if err := b.ReadQueryEntity(&req); err != nil {
		response(b, &hpb.HealthCheckResponse{}, err)
		return
	}
	var headers = make(map[string]string)
	ctx := common.NewContext(headers)
	client := hpb.NewHealthClient(ctx, archaius.GetString("servicecomb.service.name", SERVICE_NAME))
	resp, err := client.Check(&req)
	response(b, resp, err)
}

func (h *Hello) GetConfig(b *rf.Context) {
	resp := &pb.HelloResp{
		Resp: archaius.GetString("project", "xxx"),
	}
	response(b, resp, nil)
}

func (r *Hello) URLPatterns() []rf.Route {
	return []rf.Route{
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/health",
			FuncDesc:         "health",
			ResourceFuncName: "Health",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/hello",
			FuncDesc:         "say hello",
			ResourceFuncName: "Say",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/hello/grpc",
			FuncDesc:         "hello grpc",
			ResourceFuncName: "HelloGrpc",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/say",
			FuncDesc:         "query say hello",
			ResourceFuncName: "QuerySay",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/say/{say}",
			FuncDesc:         "path say hello",
			ResourceFuncName: "PathSay",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodPost,
			Path:             API_PATH + "/say",
			FuncDesc:         "body say hello",
			ResourceFuncName: "BodySay",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodPost,
			Path:             API_PATH + "/say/{say}",
			FuncDesc:         "body query path say hello",
			ResourceFuncName: "BodyQueryPathSay",
			Returns:          []*rf.Returns{{Code: 200}},
		},
		{
			Method:           http.MethodGet,
			Path:             API_PATH + "/conf",
			FuncDesc:         "get config",
			ResourceFuncName: "GetConfig",
			Returns:          []*rf.Returns{{Code: 200}},
		},
	}
}
