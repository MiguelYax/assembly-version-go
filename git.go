package main

import (
	"bytes"
	"os/exec"
	"strings"
)

var updateMsg string = "Already up to date."

func isUpToDate() bool {
	stdout, err := execute("pull")
	if err == nil && clearString(stdout) == updateMsg {
		return true
	} else {
		lg(stdout)
		return false
	}
}

func getChecksum() (string, error) {
	stdout, err := execute("log", "-n", "1", "--format=\"%h\"")
	return clearString(stdout), err
}

func commitAssemblyFile(msg string) error {
	stdout, err := execute("pull")
	if err == nil && stdout == updateMsg {
		stdout, err := execute("add", AssemblyFile)
		if err == nil {
			lg((stdout))
			stdout, err := execute("commit", "-m", "\""+msg+"\"")
			if err == nil {
				lg(stdout)
			}
		}
		// }
	}
	return err
}

func createTag(checksum string, version string, mode string) error {
	tag := "v" + version
	stdout, err := execute("tag", "-a", tag, checksum, "-m", "\""+mode+"\"")
	debugLog(stdout)
	if err == nil {

	}
	return err
}

func execute(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
