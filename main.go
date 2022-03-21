package main

import (
	_ "lczx/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"lczx/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
