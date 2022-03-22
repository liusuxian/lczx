package response

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"lczx/utility/logger"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int    `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string `json:"message"` // 提示信息
	Data    any    `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// JsonOK 标准返回结果数据结构封装。
func JsonOK(ctx context.Context, data any) {
	Json(ctx, gcode.CodeOK, data)
}

// Json 标准返回结果数据结构封装。
func Json(ctx context.Context, code gcode.Code, data any) {
	err := g.RequestFromCtx(ctx).Response.WriteJson(JsonResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    data,
	})
	if err != nil {
		logger.Error(ctx, "JsonResponse Error:", err.Error())
	}
}
