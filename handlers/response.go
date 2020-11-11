package handlers

import (
	"net/http"

	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// 错误返回
type errResp struct {
	ErrCode int    `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误描述
}

// 有消息体的错误返回
type errRespWithBody struct {
	ErrCode  int         `json:"err_code"`  // 错误码
	ErrMsg   string      `json:"err_msg"`   // 错误描述
	Success  bool        `json:"success"`   // 请求是否成功
	HttpCode int         `json:"http_code"` // http状态码
	Data     interface{} `json:"data"`      // 响应内容
}

// 统一错误处理
func response(b *rf.Context, resp interface{}, err error) {
	if err != nil {
		b.WriteHeaderAndJSON(http.StatusBadRequest,
			errResp{
				ErrCode: http.StatusBadRequest,
				ErrMsg:  err.Error(),
			},
			"application/json;charset=utf-8")
		return
	}

	b.WriteHeaderAndJSON(http.StatusOK,
		resp,
		"application/json;charset=utf-8")
	return
}
