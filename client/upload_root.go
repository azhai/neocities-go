package client

import (
	"fmt"
	"os"

	"github.com/azhai/neocities-go/api"
)

var cmdUploadRoot = &Command{
	Run:   runUploadRoot,
	Key:   "upload-root",
	Usage: "upload-root",
	Short: "Upload local files to webroot",
	Long:  "Upload all local files to webroot",
}

func init() {
	CmdRunner.Use(cmdUploadRoot)
}

func runUploadRoot(cmd *Command, args *Args) {
	cred := getCredentials()
	files := listLocalFiles()

	response, err := api.UploadFiles(cred, files)
	if err != nil {
		response.Print()

		os.Exit(1)
	}

	if isVerbose() {
		response.Print()
	}

	os.Exit(0)
}

func listLocalFiles() (files []string) {
	entries, err := os.ReadDir(".")
	if err != nil {
		if isVerbose() {
			fmt.Println(err)
		}
		return
	}

	for _, t := range entries {
		if t.Name() == ".DS_Store" {
			continue
		}
		files = append(files, t.Name())
	}
	return
}
