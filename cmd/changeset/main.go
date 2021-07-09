package main

import (
	"github.com/desmos-labs/changeset/cmd"
)

func main() {
	executor := cmd.RootCmd()
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
