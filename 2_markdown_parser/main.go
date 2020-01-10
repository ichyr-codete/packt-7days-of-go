package main

import "log"

type Tree struct {
	items Element
}

func main() {
	input := "**This is a title** Once upon a time (*long ago*) there _was_ a **_bog_** a rattling tog."

	result := ParseMarkdown(input)

	log.Printf("==================   Result  ================== \n %#v", result)
}

// MarkdownToHTML ...
func MarkdownToHTML(input string) string {
	tree := ParseMarkdown(input)
	return RenderHTML(tree)
}

// ParseMarkdown ...
func ParseMarkdown(input string) Tree {
	return Tree{}
}

// RenderHTML ...
func RenderHTML(tree Tree) string {
	return ""
}
