package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

type MarkdownFile struct {
	Title       string
	Body        string
	FrontMatter map[string]string
}

func parseMarkdown(text string) (result MarkdownFile) {
	result.FrontMatter = make(map[string]string)
	lines := strings.Split(text, "\n")
	lineNo := 0

	// Every markdown file must start with front matter.
	//
	// Each front matter line is simply a "key = value" pair, where anything
	// after the first "=" is the value. No escape mechanism here.
	//
	// There must be at least 1 empty line after front matter.
	for _, line := range lines {
		lineNo++
		key, val, found := strings.Cut(line, "=")
		if !found {
			if strings.TrimSpace(line) != "" {
				fmt.Fprintf(
					os.Stderr,
					"Line %d: Expected at least 1 empty line after front matter. Found:\n  %s\n",
					lineNo, line,
				)
				os.Exit(1)
			}
			break
		}
		key = strings.ToLower(strings.TrimSpace(key))
		val = strings.TrimSpace(val)
		if key == "title" {
			result.Title = val
		} else {
			result.FrontMatter[key] = val
		}
	}

	bodyText := strings.Join(lines[lineNo:], "\n")
	result.Body = string(markdown.ToHTML([]byte(bodyText), nil, nil))
	return
}
