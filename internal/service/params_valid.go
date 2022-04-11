package service

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
)

type sParamsValid struct{}

var (
	insParamsValid = sParamsValid{}
)

func init() {
	gvalid.RegisterRule("slice_valid", SliceValid)
}

// ParamsValid 自定义参数校验服务
func ParamsValid() *sParamsValid {
	return &insParamsValid
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
	var val int64
	for _, v := range list {
		number, ok := v.(json.Number)
		if !ok {
			err = gerror.New("不是无符号的整型切片")
			return
		}
		val, err = number.Int64()
		if err != nil || val < 0 {
			err = gerror.New("不是无符号的整型切片")
			return
		}
	}
	return
}
