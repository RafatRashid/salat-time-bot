package main

import (
	"fmt"
	"log/slog"
	"os"
	"telegram-bot/cmd"
)

func main() {
	err := cmd.Execute(os.Args[1:])
	if err != nil {
		slog.Error(fmt.Sprintf(`app exited with message: %s`, err.Error()))
	}
}
