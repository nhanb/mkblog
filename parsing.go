package main

type MarkdownFile struct {
	Title string
	Body  string
}

func parseMarkdown(mdpath string) MarkdownFile {
	// TODO: parse front matter (probably in json) and body
}
