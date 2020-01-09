package main

import (
	"reflect"
	"testing"
)

func TestParseRawInputZeroInputResult(t *testing.T) {
	a := ParseRawInput("")
	e := make([]map[string]string, 0)
	if !reflect.DeepEqual(a, e) {
		t.Error("empty string doe not result in empty map")
	}
}
func TestParseRawInputRegularSizeInput(t *testing.T) {}
func TestParseRawWrongInput(t *testing.T)            {}

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
