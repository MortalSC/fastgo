package main

import (
	"os"

	"github.com/MortalSC/fastgo/cmd/fg-apiserver/app"
)

func main() {
	command := app.NewFastGOCommand()

	if err := command.Execute(); err != nil {
		// 发生错误，则退出
		os.Exit(1)
	}
}
