package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

func mdToHtml(inpath, outpath string) {
	md, e := ioutil.ReadFile(inpath)
	panicIfErr(e)
	html := markdown.ToHTML(md, nil, nil)
	ioutil.WriteFile(outpath, html, 0644)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkblog my-blog-dir")
		os.Exit(1)
	}
	rootdir := flag.Args()[0]
	mdpaths := find(rootdir, ".md")
	for _, mdpath := range mdpaths {
		htmlpath := strings.TrimSuffix(mdpath, ".md") + ".html"
		fmt.Println(mdpath, "->", filepath.Base(htmlpath))
		mdToHtml(mdpath, htmlpath)
	}
}
