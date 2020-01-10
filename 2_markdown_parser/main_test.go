package main

import "testing"

import "reflect"

type TestData struct {
	Description string
	Input       string
	Expected    string
	Err         string
}

func TestMarkdownToHTMLInvalidMarkdownHandling(t *testing.T) {}

func TestMarkdownToHTML(t *testing.T) {
	var testData [4]TestData = [4]TestData{
		TestData{"empty input", "", "", "empty input produced some content"},
		TestData{"single italic input", "*Home*", "<i>Home</i>", "italic was not parsed"},
		TestData{"single bold  input", "**Home**", "<b>Home</b>", "bold was not parsed"},
		TestData{"single emphasized  input", "_Home_", "<em>Home</em>", "em was not parsed"},
	}

	for _, val := range testData {
		t.Run(val.Description, func(t *testing.T) {
			a := MarkdownToHTML(val.Input)
			if a != val.Expected {
				t.Error(val.Err)
			}
		})
	}
}

func TestParseMarkdown(t *testing.T) {
	var testData [4]TestData = [4]TestData{
		TestData{"text with one tag", "Hello, **user**!", , "empty input produced some content"},
	}

	for _, val := range testData {
		t.Run(val.Description, func(t *testing.T) {
			a := MarkdownToHTML(val.Input)
			if !reflect.DeepEqual(a, val.Expected) {
				t.Error(val.Err)
			}
		})
	}
}
