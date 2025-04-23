package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/azhai/neocities-go/api"
)

var cmdDeleteAll = &Command{
	Run:   runDeleteAll,
	Key:   "delete-all",
	Usage: "delete-all",
	Short: "Delete all remote files",
	Long:  "Delete all remote files except index.html",
}

func init() {
	CmdRunner.Use(cmdDeleteAll)
}

func runDeleteAll(cmd *Command, args *Args) {
	cred := getCredentials()
	files := listCurrentFiles(cred)

	response, err := api.DeleteFiles(cred, files)
	if err != nil {
		response.Print()

		os.Exit(1)
	}

	if isVerbose() {
		response.Print()
	}

	os.Exit(0)
}

func listCurrentFiles(cred api.Credentials) (files []string) {
	list, err := api.List(cred)
	if err != nil {
		if isVerbose() {
			fmt.Println(err)
		}
		return
	}

	lastDir := "/"
	for _, f := range list.Files {
		if f.Path == "index.html" {
			continue
		} else if strings.HasPrefix(f.Path, lastDir) {
			continue
		}
		if f.IsDirectory {
			lastDir = f.Path + "/"
		}
		files = append(files, f.Path)
	}
	return
}
