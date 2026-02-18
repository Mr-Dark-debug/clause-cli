// Package main is the entry point for the Clause CLI.
package main

import (
	"os"

	"github.com/clause-cli/clause/internal/cmd"
)

func main() {
	os.Exit(cmd.ExecuteWithError())
}
