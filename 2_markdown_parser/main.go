package main

import "log"

func main() {
	input := "**This is a title** Once upon a time (*long ago*) there _was_ a **_bog_** a rattling tog."

	result := ParseMarkdown(input)
	log.Printf("==================   Result  ================== \n %#v", result)
}

// ParseMarkdown ...
func ParseMarkdown(input string) string {
	return ""
}
