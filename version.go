package main


import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

const (
	MAJOR string = "major"
	MINOR string = "minor"
	BUILD string = "build"
	REVISION string = "revision"
)


func parse(currentVersion string) (int, int, int, int)  {
		items := strings.Split(currentVersion, ".")
		major, err1 := strconv.Atoi(items[0])
		minor, err2 := strconv.Atoi(items[1])
		build, err3 := strconv.Atoi(items[2])
		revision, err4 := strconv.Atoi(items[3])
		if (err1 != nil || err2 != nil || err3 != nil || err4 != nil ) {
			return 0,0,0,0
		} else {
			return major,minor,build,revision
		}
}

func isValid(currentVersion string)  (bool, error) {
	matched, err := regexp.MatchString(`\d+\.\d+\.\d+\.\d+`, currentVersion)
	if (matched && err == nil ) {
			return true, err
		} else {
			return false, err
		}

}
//Up change version based on mode parameter 
func Up(tag string, mode string) (string, error) {
 	return	change(tag,mode,true)
}

func down(tag string, mode string) (string, error) {
	return change(tag,mode,false)
}

func change(tag string, mode string, increment bool) (string, error)  {
	var newTag string = ""
	valid, err := isValid(tag)
	if (valid) {
		var add int 
		if (increment) {
			add = 1
		} else {
			add = -1
		}
		major, minor, build, revision := parse(tag)
	
		switch mode {
		case MAJOR:
			major = major + add
			minor = 0
			build = 0
			revision = 0
		case MINOR:
			minor += add
			build = 0
			revision = 0
		case BUILD:
			build += add
			revision = 0
		case REVISION:
			revision += add
		}
		newTag = fmt.Sprintf("%d.%d.%d.%d", major,minor,build,revision)	
	}
	return newTag, err
}