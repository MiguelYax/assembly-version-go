package main

import (
	"io/ioutil"
	"regexp"
	"strings"

)

var AssemblyVersionTag string
var AssemblyFileVersionTag string
var Content string
var AssemblyFile string = "AssemblyInfo.cs"

func getCurrentTag() (string, string, error) {
	data, err := ioutil.ReadFile(AssemblyFile)
	if (err == nil) { 
		Content = string(data)
		assemblyVersionRegexp := regexp.MustCompile(`AssemblyVersion\("\d+\.\d+\.\d+\.\d+"\)`)
		assemblyFileVersionRegexp := regexp.MustCompile(`AssemblyFileVersion\("\d+\.\d+\.\d+\.\d+"\)`)
		AssemblyVersionTag = assemblyVersionRegexp.FindString(Content)
		AssemblyFileVersionTag = assemblyFileVersionRegexp.FindString(Content)
		versionRegex := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
		assemblyVersion := versionRegex.FindString(AssemblyVersionTag)
		assemblyFileVersion := versionRegex.FindString(AssemblyFileVersionTag)
		return assemblyVersion,assemblyFileVersion, nil
	} else {
		return "", "", err
	}
}

func writeCurrentTag(assemblyVersion string, assemblyFileVersion string, all bool) ( error) {
		var newAssemblyVersionTag = "AssemblyVersion(\""+ assemblyVersion + "\")"
		var newAssemblyFileVersionTag = "AssemblyFileVersion(\""+ assemblyFileVersion + "\")"
		 newContent := strings.Replace(Content, AssemblyVersionTag, newAssemblyVersionTag,1)
		 newContent = strings.Replace(newContent, AssemblyFileVersionTag, newAssemblyFileVersionTag,1)

		data := []byte(newContent)
	    err := ioutil.WriteFile(AssemblyFile, data, 0644)

	return  err
}