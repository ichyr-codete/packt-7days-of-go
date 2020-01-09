package main

import "testing"

import "reflect"

func TestFilterEmptyStringsResultingInEmptyArray(t *testing.T) {
	result := FilterEmptyStrings([]string{"", "", ""})
	if len(result) > 0 {
		t.Error("didn't receive empty array")
	}
}

func TestFilterEmptyStrings(t *testing.T) {
	expected := []string{"a"}
	actual := FilterEmptyStrings([]string{"", "a", ""})
	if !reflect.DeepEqual(expected, actual) {
		t.Error("resulting slices are not equal")
	}
}
