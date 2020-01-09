package main

import (
	"reflect"
	"testing"
)

func TestParseRawInputZeroInputResult(t *testing.T) {
	a, _ := ParseRawInput("")
	e := make([]map[string]string, 0)
	if !reflect.DeepEqual(a, e) {
		t.Error("empty string doe not result in empty map")
	}
}

func TestParseRawInputOneLineInput(t *testing.T) {
	a, _ := ParseRawInput(`
	UID   PID  PPID   C STIME   TTY           TIME CMD               %CPU
	501  4248  3893   0  1:11pm ttys002    0:00.25 go run main.go    14.2  501  4263  4248   0  1:11pm ttys002    0:00.00 /var/folders/5c/   0.0
	`)
	e := []map[string]string{
		map[string]string{"%CPU": "26.6", "C": "0", "CMD": "go run main.go", "PID": "5351", "PPID": "3893", "STIME": "1:29pm", "TIME": "0:00.27", "TTY": "ttys002", "UID": "501"},
	}
	if reflect.DeepEqual(a, e) {
		t.Error("failed to parse one line")
	}
}

func TestParseRawInputRegularSizeInput(t *testing.T) {
	a, _ := ParseRawInput(`
	
	UID   PID  PPID   C STIME   TTY           TIME CMD               %CPU
	501  4248  3893   0  1:11pm ttys002    0:00.25 go run main.go    14.2  501  4263  4248   0  1:11pm ttys002    0:00.00 /var/folders/5c/   0.0
	501  3893  2946   0  1:11pm ttys002    0:00.05 /bin/bash -l       0.0
	501   505   480   0 10:51am ttys000    0:00.06 -bash              0.0
	501   485   483   0 10:51am ttys001    0:00.06 -bash              0.0
	
	`)
	if len(a) != 4 {
		t.Errorf("Resulting slice should have 4 entries, but has %v", len(a))
	}
}
func TestParseRawWrongInput(t *testing.T) {
	_, err := ParseRawInput(`
	PID TTY           TIME CMD
    1 ??         3:50.91 /sbin/launchd
  132 ??         0:13.35 /usr/sbin/syslogd
  133 ??         0:04.21 /usr/libexec/UserEventAgent (System)
	`)
	if err.Error() != "ps command resulted in different data" {
		t.Errorf("wrong input did not trigger error to be returned")
	}
}

func TestFilterEmptyStringsResultingInEmptyArray(t *testing.T) {
	result := FilterEmptyStrings([]string{"", "", ""})
	if len(result) > 0 {
		t.Error("didn't receive empty array")
	}
}

func TestFilterEmptyStringsExtracsSingle(t *testing.T) {
	t.Run("one element result", func(t *testing.T) {
		expected := []string{"a"}
		actual := FilterEmptyStrings([]string{"", "a", ""})
		if !reflect.DeepEqual(expected, actual) {
			t.Error("resulting slices are not equal")
		}
	})
}
