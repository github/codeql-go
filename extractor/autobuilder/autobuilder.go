package autobuilder

import (
	"log"
	"os"
	"os/exec"

	"github.com/github/codeql-go/extractor/util"
)

var CheckExtracted = false

func checkEmpty(dir string) (bool, error) {
	if !util.DirExists(dir) {
		return true, nil
	}

	d, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return false, err
	}
	return len(names) == 0, nil
}

func checkExtractorRun() bool {
	srcDir := os.Getenv("CODEQL_EXTRACTOR_GO_SOURCE_ARCHIVE_DIR")
	if srcDir != "" {
		empty, err := checkEmpty(srcDir)
		if err != nil {
			log.Fatalf("Unable to read source archive directory %s.", srcDir)
		}
		if empty {
			log.Printf("No Go code seen; continuing to try other builds.")
			return false
		}
		return true
	} else {
		log.Fatalf("No source directory set.")
		return false
	}
}

func tryBuildIfExists(buildFile, cmd string, args ...string) bool {
	if util.FileExists(buildFile) {
		log.Printf("%s found, running %s\n", buildFile, cmd)
		tryBuild(cmd, args...)
	}
	return false
}

func tryBuild(cmd string, args ...string) bool {
	res := util.RunCmd(exec.Command(cmd, args...))
	return res && (!CheckExtracted || checkExtractorRun())
}

func Autobuild() bool {
	// if there is a build file, run the corresponding build tool
	return tryBuildIfExists("Makefile", "make") ||
		tryBuildIfExists("makefile", "make") ||
		tryBuildIfExists("GNUmakefile", "make") ||
		tryBuildIfExists("build.ninja", "ninja") ||
		tryBuildIfExists("build", "./build") ||
		tryBuildIfExists("build.sh", "./build.sh") ||
		tryBuild("go", "build", "./...")
}
