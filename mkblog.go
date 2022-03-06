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

	"github.com/gomarkdown/markdown"
)

func panicIfErr(e error) {
	// Gonna default to panicking at every error until I finally bother to
	// learn The Right Way (tm) to juggle them in Go.
	if e != nil {
		panic(e)
	}
}

func find(root, ext string) []string {
	var results []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		panicIfErr(e)
		if filepath.Ext(d.Name()) == ext {
			results = append(results, s)
		}
		return nil
	})
	return results
}

type Page struct {
	Title string
	Body  string
}

func mdToHtml(rootpath, inpath, outpath string) {
	md, err := ioutil.ReadFile(inpath)
	panicIfErr(err)
	bodyHtml := string(markdown.ToHTML(md, nil, nil))

	templatePath := rootpath + "_templates/base.html"
	tmpl := template.Must(template.ParseFiles(templatePath))

	outfile, err := os.Create(outpath)
	panicIfErr(err)
	defer outfile.Close()
	tmpl.Execute(outfile, Page{Title: "MkBlog", Body: bodyHtml})
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkblog my-blog-dir")
		os.Exit(1)
	}
	rootpath := flag.Args()[0]
	mdpaths := find(rootpath, ".md")
	for _, mdpath := range mdpaths {
		htmlpath := strings.TrimSuffix(mdpath, ".md") + ".html"
		fmt.Println(mdpath, "->", filepath.Base(htmlpath))
		mdToHtml(rootpath, mdpath, htmlpath)
	}
}
