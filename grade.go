package main

import (
	"os"

	"github.com/sv-tools/grade/cmd"
)

var version = "v0.0.0"

func main() {
	cmd.RootCmd.Version = version
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
