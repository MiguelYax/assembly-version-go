package main
import (
	// "bytes"	
	// "fmt"
	"log"
	"os"
	"strings"
)


func main ()  {
	if ( len(os.Args) >= 1) {
		args:= os.Args
		var mode string =  strings.ToLower(args[0])
		mode= "build"
		if (mode == MAJOR  || mode == MINOR || mode == BUILD || mode == REVISION) {
			lg("Sarting mode:" + mode)
			assemblyVersion, assemblyFileVersion, err := getCurrentTag()
			check(err)
			newAssemblyVersion, err := Up(assemblyVersion, mode)
			check(err)
			newAssemblyFileVersion, err := Up(assemblyFileVersion,mode)
			check(err)
			lg("Change assembly version: " + assemblyVersion + " to " + newAssemblyVersion)
			lg("Change assembly file version: " + assemblyFileVersion + " to " + newAssemblyFileVersion)
			
		} else {
			lg("supported modes are: major, minor, build, revision ")
		}
	} else {
		lg("show help -h --help")
	}
}

func check(e error) {
    if e != nil {
        log.Panic(e)
    }
}

func lg(msg string)  {
	log.Print(msg)
}