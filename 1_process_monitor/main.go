package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// query os
	out, err := ExecutePS("-fr", "-o pcpu")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" ========= Command line output ========= ")
	rawInput := out.String()
	log.Println("\n", rawInput)

	// strip first line
	// parse all data as map
	data, err := ParseRawInput(rawInput)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(" ========= Parsed data table ========= ")
	log.Printf("%#v", data)

	// present cpu intensive process

}

// ExecutePS ...
func ExecutePS(args ...string) (*bytes.Buffer, error) {
	var cmd *exec.Cmd
	if len(args) > 0 {
		cmd = exec.Command("ps", args...)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return &out, err
}

// ParseRawInput ...
func ParseRawInput(input string) ([]map[string]string, error) {
	headerReg := regexp.MustCompile(`UID\s+PID\s+PPID\s+C\s+STIME\s+TTY\s+TIME\s+CMD\s+%CPU`)

	reg := regexp.MustCompile(`(?P<uid>\d+)\s+(?P<pid>\d+)\s+(?P<ppid>\d+)\s+(?P<c>\d+)\s+(?P<stime>\d?\d:\d{2}\w{2})\s+(?P<tty>\w+)\s+(?P<time>\d+:\d+.\d+)\s+(?P<cmd>.+)\s+(?P<percent_cpu>\d+.\d+)`)

	rows := strings.Split(input, "\n")

	if len(rows) == 1 {
		return make([]map[string]string, 0), nil
	}

	if headerReg.FindString(rows[0]) == "" {
		return make([]map[string]string, 0), fmt.Errorf("ps command resulted in different data")
	}

	rawHeaders := strings.Split(rows[0], " ")
	headers := FilterEmptyStrings(rawHeaders)
	dataLength := len(headers)
	dataRows := rows[1:]
	result := make([]map[string]string, len(dataRows))
	for _, value := range dataRows {
		if value == "" {
			continue
		}

		// These routines take an extra integer argument, n.
		// If n >= 0, the function returns at most n matches/submatches;
		// otherwise, it returns all of them.
		data := reg.FindAllStringSubmatch(value, -1)

		row := make(map[string]string, dataLength)
		for index, val := range data[0] {
			if index == 0 {
				continue
			}
			row[headers[index-1]] = strings.TrimSpace(val)
		}
		result = append(result, row)
	}
	return result, nil
}

// FilterEmptyStrings ...
func FilterEmptyStrings(arr []string) []string {
	var res = []string{}
	for _, val := range arr {
		if val == "" {
			continue
		}
		res = append(res, val)
	}
	return res
}
