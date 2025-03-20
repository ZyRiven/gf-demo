package main

import (
	_ "sleep-service/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"sleep-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
