package main

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
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

	// strip first line
	// parse all data as map
	data := parseRawInput(rawInput)
	log.Println(" ========= Parsed data table ========= ")
	log.Printf("%#v", data)
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

func parseRawInput(input string) []map[string]string {
	reg := regexp.MustCompile(`(?P<uid>\d+)\s+(?P<pid>\d+)\s+(?P<ppid>\d+)\s+(?P<c>\d+)\s+(?P<stime>\d?\d:\d\dpm)\s+(?P<tty>\w+)\s+(?P<time>\d+:\d+.\d+)\s+(?P<cmd>.+)(?P<percent_cpu>\d+.\d+)`)

	rows := strings.Split(input, "\n")
	rawHeaders := strings.Split(rows[0], " ")
	headers := filterEmptyStrings(rawHeaders)
	var result []map[string]string
	dataRows := rows[1:]
	for index, value := range dataRows {
		// These routines take an extra integer argument, n.
		// If n >= 0, the function returns at most n matches/submatches;
		// otherwise, it returns all of them.
		data := reg.FindAllStringSubmatch(value, -1)
		row := make(map[string]string, len(headers))
		count := 0
		for _, val := range data {
			// if val == []string{""} {
			// 	continue
			// }
			row[headers[count]] = val[0]
			count++
		}
		result[index] = row
	}
	return result
}

func filterEmptyStrings(arr []string) []string {
	var res = []string{}
	for _, val := range arr {
		if val == "" {
			continue
		}
		res = append(res, val)
	}
	return res
}
