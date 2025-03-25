package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "sleep-service/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"sleep-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
