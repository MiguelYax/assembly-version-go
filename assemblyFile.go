package main

import (
	"io/ioutil"
	"regexp"
)

func getCurrentTag() (string, string, error) {
	data, err := ioutil.ReadFile("AssemblyInfo.cs")
	if (err == nil) { 
		content := string(data)
		assemblyVersionRegexp := regexp.MustCompile(`AssemblyVersion\("\d+\.\d+\.\d+\.\d+"\)`)
		assemblyFileVersionRegexp := regexp.MustCompile(`AssemblyVersion\("\d+\.\d+\.\d+\.\d+"\)`)
		assemblyVersionTag := assemblyVersionRegexp.FindString(content)
		assemblyFileVersionTag := assemblyFileVersionRegexp.FindString(content)
		versionRegex := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
		assemblyVersion := versionRegex.FindString(assemblyVersionTag)
		assemblyFileVersion := versionRegex.FindString(assemblyFileVersionTag)
		return assemblyVersion,assemblyFileVersion, nil
	} else {
		return "", "", err
	}
}