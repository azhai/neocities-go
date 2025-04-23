/*
A Neocities API client written in Go

# Installation

Just go get the command:

	go get -u github.com/azhai/neocities-go

# Usage

First you need to export some environment variables:

	export NEOCITIES_USER=<username>
	export NEOCITIES_PASS=<password>

Then you can run the command:

	neocities upload <filename> [<another filename>]
*/
package main

import (
	"os"

	"github.com/azhai/neocities-go/client"
)

func main() {
	err := client.CmdRunner.Execute()

	os.Exit(err.ExitCode)
}
