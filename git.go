package main

import (
	"bytes"	
	"fmt"
	"log"
	"os/exec"
	"strings"
)
func gitPull()  {

	println("Start")
	cmd := exec.Command("git", "pull")	
	cmd.Stdin = strings.NewReader("Some imput")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("in all caps: %q\n", out.String())
	
}