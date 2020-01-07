package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	// query os
	out, err := executePS("-fr", "-o pcpu")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" ========= Command line output ========= ")
	rawInput := out.String()
	log.Println("\n", rawInput)
}

func executePS(args ...string) (*bytes.Buffer, error) {
	var cmd *exec.Cmd
	if len(args) > 0 {
		cmd = exec.Command("ps", args...)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return &out, err
}
