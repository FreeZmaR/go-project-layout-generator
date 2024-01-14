package main

import (
	"github.com/FreeZmaR/go-project-layout-generator/terminal"
	"log/slog"
	"os"
)

func main() {
	term := terminal.New()

	if err := term.Run(); err != nil {
		slog.Error("error running terminal: ", slog.String("error", err.Error()))

		os.Exit(1)
	}
}
