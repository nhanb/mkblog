package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func panicIfErr(e error) {
	// Gonna default to panicking at every error until I finally bother to
	// learn The Right Way (tm) to juggle them in Go.
	if e != nil {
		panic(e)
	}
}

func find(ext string) []string {
	var results []string
	filepath.WalkDir(".", func(s string, d fs.DirEntry, e error) error {
		panicIfErr(e)
		if filepath.Ext(d.Name()) == ext {
			results = append(results, s)
		}
		return nil
	})
	return results
}

type Page struct {
	AutogenWarning string
	Title          string
	Body           string
}

func mdToHtml(inpath string) {
	md, err := ioutil.ReadFile(inpath)
	panicIfErr(err)
	parsed := parseMarkdown(string(md))

	outpath := strings.TrimSuffix(inpath, ".md") + ".html"
	fmt.Println(inpath, "->", filepath.Base(outpath))

	templatePath := filepath.Join("_templates", "base.html")
	tmpl := template.Must(template.ParseFiles(templatePath))

	outfile, err := os.Create(outpath)
	panicIfErr(err)
	defer outfile.Close()

	tmpl.Execute(
		outfile,
		Page{
			AutogenWarning: "This file was auto-generated by mkblog. Do not edit directly.",
			Title:          parsed.Title,
			Body:           parsed.Body,
		},
	)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkblog my-blog-dir")
		os.Exit(1)
	}
	blogPath := filepath.Clean(flag.Args()[0])
	os.Chdir(blogPath)
	fmt.Println("blogPath:", blogPath)

	mdpaths := find(".md")
	for _, mdpath := range mdpaths {
		mdToHtml(mdpath)
	}
}
