package main

import (
	"fmt"
	"os"

	"github.com/spilliams/foo/internal/cli"
)

func main() {
	err := cli.NewFooCmd().Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
