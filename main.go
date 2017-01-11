package main

import (
	"os"

	"gitlab.com/SporeDB/sporedb/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
