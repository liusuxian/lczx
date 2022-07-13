package params_valid

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
	"lczx/internal/service"
	"strconv"
)

type sParamsValid struct{}

func init() {
	service.RegisterParamsValid(newParamsValid())
}

// 自定义参数校验服务
func newParamsValid() *sParamsValid {
	return &sParamsValid{}
}

// RegisterRule 注册参数校验规则
func (s *sParamsValid) RegisterRule() {
	gvalid.RegisterRule("slice_valid", SliceValid)
}

// SliceValid 自定义切片数据校验
func SliceValid(ctx context.Context, in gvalid.RuleFuncInput) (err error) {
	rules := gstr.Split(in.Rule, ":")
	if len(rules) < 2 {
		err = gerror.Newf("无效的自定义规则: %s", in.Rule)
		return
	}

	list := in.Value.Interfaces()
	switch rules[1] {
	case "uint64", "uint":
		err = uintSliceValid(list)
	default:
		err = gerror.Newf("无效的自定义规则: %s", in.Rule)
	}
	return
}

// 无符号整数切片验证
func uintSliceValid(list []any) (err error) {
	for _, v := range list {
		switch value := v.(type) {
		case json.Number:
			var val int64
			val, err = value.Int64()
			if err != nil || val < 0 {
				err = gerror.New("不是无符号的整型切片")
				return
			}
		case string:
			var val int
			val, err = strconv.Atoi(value)
			if err != nil || val < 0 {
				err = gerror.New("不是无符号的整型切片")
				return
			}
		default:
			err = gerror.New("不是无符号的整型切片")
			return
		}
	}
	return
}
