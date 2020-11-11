package handlers

import (
	"net/http"

	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

func init() {
	chassis.RegisterSchema("rest", &Server{}, server.WithSchemaID(SERVICE_NAME))
}

type Api struct {
	ApiVersion string `json:"api_version"`
}

type Server struct{}

func (h *Server) Version(b *rf.Context) {
	response(b,
		&Api{
			ApiVersion: API_VERSION,
		},
		nil)
}

func (r *Server) URLPatterns() []rf.Route {
	return []rf.Route{
		{
			Method:           http.MethodGet,
			Path:             "/",
			FuncDesc:         "system version",
			ResourceFuncName: "Version",
			Returns:          []*rf.Returns{{Code: 200}},
		},
	}
}
