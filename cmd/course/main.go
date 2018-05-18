package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const pptBasics = "tmpl/ppt.html"
const addr = ":3000"

func main() {
	http.HandleFunc("/", serveGoBasicsPresentation)
	fmt.Printf("Serving %q on %q...", pptBasics, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func serveGoBasicsPresentation(w http.ResponseWriter, r *http.Request) {
	tmpl, err := loadPresentationTemplate(pptBasics)
	if err != nil {
		log.Panicf("Executing %q: %v", pptBasics, err)
	}
	tmpl.Execute(w, pptBasics)
}

func loadPresentationTemplate(ppt string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(ppt)
	if err != nil {
		return nil, fmt.Errorf("Loading ppt from: %q: %v", ppt, err)
	}
	dir := "samples"
	if err := addChildTemplates(tmpl, dir, ".go"); err != nil {
		return nil, fmt.Errorf("Loading template from: %q: %v", dir, err)
	}
	fmt.Printf("Templates Loaded%s\n", tmpl.DefinedTemplates())
	return tmpl, nil
}

func addChildTemplates(parent *template.Template, dir string, ext string) error {
	gofiles, err := listFiles(dir, ext)
	if err != nil {
		return fmt.Errorf("Error listing %q files: %v", ext, err)
	}
	if len(gofiles) == 0 {
		log.Panic("No go files found!")
	}
	for _, f := range gofiles {
		sample, err := ioutil.ReadFile(f)
		if err != nil {
			return fmt.Errorf("Cannot read go file %q: %v", f, err)
		}
		n := strings.TrimPrefix(f, "samples/")
		t, err := template.New(n).Parse(string(sample))
		if err != nil {
			return fmt.Errorf("Parsing %q: %v", f, err)
		}
		parent.AddParseTree(n, t.Tree)
	}
	return nil
}

func listFiles(dir string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Walking dir %s: %v", dir, err)
	}
	return files, err
}
