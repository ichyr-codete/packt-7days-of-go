package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// query os
	out, err := executePS("-p 61960")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

func executePS(args ...string) (*bytes.Buffer, error) {
	var cmd *exec.Cmd
	if len(args) > 0 {
		cmd = exec.Command("ps", args...)
	}
	log.Printf("%#v\n", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return &out, err
}
