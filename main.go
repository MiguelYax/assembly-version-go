package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 1 {
		args := os.Args
		var mode string = strings.ToLower(args[0])
		mode = "build"
		all := true
		if mode == MAJOR || mode == MINOR || mode == BUILD || mode == REVISION {
			if isUpToDate() {
				lg("Sarting mode:" + mode)
				assemblyVersion, assemblyFileVersion, err := getCurrentTag()
				check(err)
				newAssemblyVersion, err := Up(assemblyVersion, mode)
				check(err)
				newAssemblyFileVersion, err := Up(assemblyFileVersion, mode)
				check(err)
				err = writeCurrentTag(newAssemblyVersion, newAssemblyFileVersion, all)
				versionMsg := " Version: " + assemblyVersion + " to " + newAssemblyVersion
				fileVersionMsg := " File version: " + assemblyFileVersion + " to " + newAssemblyFileVersion
				lg(versionMsg)
				lg(fileVersionMsg)
				if err == nil {
					if newAssemblyVersion == newAssemblyFileVersion {
						commitAssemblyFile(versionMsg)
					} else {
						commitAssemblyFile(versionMsg + fileVersionMsg)
					}
				}
				// commitAssemblyFile()
			}

		} else {
			lg("supported modes are: major, minor, build, revision ")
		}
	} else {
		lg("show help -h --help")
	}
}
