package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "lczx/internal/packed"

	_ "lczx/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"lczx/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
