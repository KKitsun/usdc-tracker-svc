package main

import (
	"os"

	"github.com/KKitsun/usdc-tracker-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
