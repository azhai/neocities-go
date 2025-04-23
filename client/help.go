package client

import (
	"fmt"
	"os"
)

var cmdHelp = &Command{
	Run:   runHelp,
	Usage: "help [command]",
	Short: "Show help",
	Long:  "Show usage instructions for a command",
}

func init() {
	CmdRunner.Use(cmdHelp)
}

func runHelp(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		printUsage()

		os.Exit(0)
	}

	for _, cmd := range CmdRunner.All() {
		if cmd.Name() == args.FirstParam() {
			cmd.PrintUsage()

			os.Exit(0)
		}
	}
}

var helpText = `usage: neocities <command> [<args>]

Commands:
   upload       Upload files to Neocities
   upload-root  Upload local files to webroot
   delete       Delete files from Neocities
   delete-all   Delete all remote files
   info         Info about Neocities websites
   key          Neocities API key
   list         List files on Neocities
   version      Show neocities client version

Help for a specific command:
   help [command]

Environment setup:

   export NEOCITIES_USER=<username>
   export NEOCITIES_PASS=<password>

  (OR)

   export NEOCITIES_API_KEY=<key>

`

func printUsage() {
	fmt.Print(helpText)
}
