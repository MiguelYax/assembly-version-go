package main

import (
	"log"
	"strings"
)

var debug bool = true

func check(e error) {
	if e != nil {
		if debug {
			log.Panic(e)
		} else {
			log.Print(e)
		}
	}
}

func lg(msg string) {
	log.Print(msg)
}

func debugLog(msg string) {
	if debug {
		lg(msg)
	}
}

func clearString(str string) string {
	newStr := strings.Trim(strings.Replace(str, "\n", "", -1), "")

	return newStr
}
