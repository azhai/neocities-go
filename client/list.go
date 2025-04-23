package client

import (
	"fmt"

	"github.com/azhai/neocities-go/api"
)

var cmdList = &Command{
	Run:   runList,
	Usage: "list",
	Short: "List files on Neocities",
	Long:  "List files in your Neocities website",
}

func init() {
	CmdRunner.Use(cmdList)
}

func runList(cmd *Command, args *Args) {
	cred := getCredentials()

	list, err := api.List(cred)
	if err != nil {
		fmt.Println(err)
		return
	}

	dump(list)
}
