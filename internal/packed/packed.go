package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
