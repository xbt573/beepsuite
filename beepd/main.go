package main

import (
	"log/slog"

	"github.com/xbt573/beepsuite/beepd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("failed to run cmd", "err", err)
	}
}
