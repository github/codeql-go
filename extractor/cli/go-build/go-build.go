package main

import (
	"github.com/github/codeql-go/extractor/util"
	"os/exec"

	"github.com/github/codeql-go/extractor/autobuilder"
)

func main() {
	// check if a build command has successfully extracted something
	autobuilder.CheckExtracted = true

	autobuilder.Autobuild()

	util.RunCmd(exec.Command("go", "build", "./..."))
}
