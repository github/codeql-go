package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/codeskyblue/go-sh"
	"golang.org/x/crypto/ssh"
)

func main() {}

func handler(w http.ResponseWriter, req *http.Request) {
	sudo := "sudo"
	shell := "/bin/bash"
	assumedNonShell := "ls"
	args := []string{}

	source := req.URL.Query()["cmd"][0]

	// os.StartProcess: these MUST be caught.
	{
		// `source` is used directly as command:
		os.StartProcess(source, args, nil)

		// `source` flows into a composite literal which is used
		// as arguments, and the command is a shell:
		os.StartProcess(shell, []string{source}, nil)

		// `source` flows into a composite literal as first argument to append:
		os.StartProcess(shell, append([]string{source}, args...), nil)

		// `source` flows into a composite literal as Nth argument to append:
		os.StartProcess(shell, append([]string{sudo}, source), nil)

		// TODO:
		//shellToAssumedNonShell := append([]string{sudo}, shell, source)
		//os.StartProcess(assumedNonShell, shellToAssumedNonShell, nil)
	}

	// os.StartProcess: these MUST NOT be caught because non-valid.
	{
		// `source` is an argument to a non-shell command that does not execute
		// the `source` as a command, i.e. the source is just an argument to a command
		// that will not execute it.
		os.StartProcess(assumedNonShell, []string{source}, nil)

		// as above, except the source is inside a composite literal inside an append:
		os.StartProcess(assumedNonShell, append([]string{source}, args...), nil)

		// source is used inside append as nth argument:
		os.StartProcess(assumedNonShell, append([]string{assumedNonShell}, source), nil)
	}

	// exec.Command: these MUST be caught.
	{
		// source is used directly as command:
		exec.Command(source, args...).Run()

		// source comes as nth arg to a shell:
		exec.Command(shell, "a0", "a1", source)

		// source flows into a composite literal used as args, and the command is a shell:
		exec.Command(shell, []string{"a0", "a1", source}...)

		// source flows into a composite literal as first argument to append:
		exec.Command(shell, append([]string{source}, args...)...)

		// source flows into a composite literal as Nth argument to append:
		exec.Command(shell, append([]string{sudo}, source)...)

		// other ways to compose a command:
		exec.Command("sh", "-c", source)
		exec.Command("sudo", "sh", "-c", source)

		// (this is a shell, so you can just create a pipe with `|` or a `$(command)` or a `&& command`):
		{
			// string concatenation:
			exec.Command("sh", "-c", "GOOS=windows GOARCH=386 go build -ldflags \"-s -w -H=windowsgui\" -o \""+source+".go")

			// source is argument to fmt.Sprintf:
			exec.Command("/bin/bash", "-c", fmt.Sprintf("ls %s", source))

			// source passes through append and strings.Join:
			exec.Command("sh", "-c", strings.Join(append([]string{assumedNonShell}, source), " "))
			exec.Command("sh", "-c", strings.Join(append([]string{assumedNonShell}, "a0", "a1", source), " "))

			// source passes through composite literal, append and strings.Join:
			exec.Command("sh", "-c", assumedNonShell, strings.Join(append([]string{source}, "a0", "a1"), " "))

			// source passes through path.Join:
			exec.Command("/bin/sh", path.Join("scripts", source))
		}
	}
	// golang.org/x/crypto/ssh
	{
		session := &ssh.Session{}
		session.CombinedOutput(source)
		session.Output(source)
		session.Run(source)
		session.Start(source)
	}
	// github.com/codeskyblue/go-sh
	{
		sh.Command(shell, toInterfaceArray(append([]string{assumedNonShell}, source)...)...)
		sh.InteractiveSession().Call(shell, toInterfaceArray(append([]string{assumedNonShell}, source)...)...)
		sh.InteractiveSession().Command(shell, toInterfaceArray(append([]string{assumedNonShell}, source)...)...)
	}
}
func toInterfaceArray(str ...string) []interface{} {
	res := make([]interface{}, 0)
	for i := range str {
		res = append(res, str[i])
	}
	return res
}
