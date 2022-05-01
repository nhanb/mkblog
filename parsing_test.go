package main

import (
	"reflect"
	"testing"
)

func TestParseFrontMatter(t *testing.T) {
	sample := `title = My name is Foo
foo =hey now
BAR= BAR should be bar

## heading

body
`
	got := parseMarkdown(sample)
	want := MarkdownFile{
		Title: "My name is Foo",
		Body:  "<h2>heading</h2>\n\n<p>body</p>\n",
		FrontMatter: map[string]string{
			"foo": "hey now",
			"bar": "BAR should be bar",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseMarkdown() = %q, want %q", got, want)
	}
}
