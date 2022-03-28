package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"lczx/utility/logger"
)

const SUCCESS = 0

// Resp 数据返回通用JSON数据结构
type Resp struct {
	Code int    `json:"code"` // 错误码(0:成功, >=1:错误码)
	Msg  string `json:"msg"`  // 提示信息
	Data any    `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// RespJson 标准返回结果数据
func RespJson(req *ghttp.Request, code int, msg string, data ...any) {
	var respData any
	if len(data) > 0 {
		respData = data[0]
	}
	req.SetParam("apiReturnResp", respData)
	err := req.Response.WriteJson(&Resp{
		Code: code,
		Msg:  msg,
		Data: respData,
	})
	if err != nil {
		logger.Error(req.GetCtx(), "RespJson Error: ", err.Error())
	}
}

// RespJsonExit 标准返回结果数据并退出
func RespJsonExit(req *ghttp.Request, code int, msg string, data ...any) {
	RespJson(req, code, msg, data...)
	req.Exit()
}

// Succ 成功
func Succ(req *ghttp.Request, data ...any) {
	RespJson(req, SUCCESS, "success", data...)
}

// SuccExit 成功并退出
func SuccExit(req *ghttp.Request, data ...any) {
	RespJsonExit(req, SUCCESS, "success", data...)
}
