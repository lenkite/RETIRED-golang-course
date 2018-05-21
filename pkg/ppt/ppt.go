package ppt

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// A Config represents the configuration for a Presentation
type Config struct {
	// Name is the name of the presentation
	Name string

	// Path represents the path to the primary html template file of the presentation
	Path string

	// RefExt is the file extension of the referenced child templates
	RefExt string

	// RefDir is the root directory under which referenced child templates are kept
	RefDir string

	// HotReload indicates whether the Presentation should hot-reload templates on change
	HotReload bool
}

// A Presentation represenets a web-based presentation that can be written to a io.Writer
type Presentation struct {
	Config

	// root is the html template of the presentation
	root *template.Template

	// refPaths is a map of child template paths embedded in the presentation to their last modified times
	refPaths map[string]time.Time
}

// New returns a new Presentation for the given configuration
// All templates are loaded and the presentation can be written to an io.Writer
func New(config Config) (*Presentation, error) {
	tmpl, err := template.ParseFiles(config.Path)
	if err != nil {
		return nil, fmt.Errorf("Loading ppt from: %q: %v", config.Path, err)
	}
	if !strings.HasPrefix(config.RefExt, ".") {
		config.RefExt = "." + config.RefExt
	}
	refPaths, err := addRefTemplates(tmpl, config.RefDir, config.RefExt)
	if err != nil {
		return nil, fmt.Errorf("Loading templates from: %q: %v", config.RefDir, err)
	}
	return &Presentation{Config: config, root: tmpl, refPaths: refPaths}, nil
}

func (ppt *Presentation) WriteTo(w io.Writer) (n int64, err error) {
	var buf bytes.Buffer
	ppt.root.Execute(&buf, ppt.Name)
	return buf.WriteTo(w)
}

func addRefTemplates(parent *template.Template, dir string, ext string) (map[string]time.Time, error) {
	refPaths := make(map[string]time.Time)
	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			refPaths[path] = fi.ModTime()
		}
		return addRefTemplate(parent, path)
	})
	if err != nil {
		return nil, fmt.Errorf("Walking dir %s: %v", dir, err)
	}
	return refPaths, nil
}

func addRefTemplate(parent *template.Template, refPath string) error {
	sample, err := ioutil.ReadFile(refPath)
	if err != nil {
		return fmt.Errorf("Cannot read go file %q: %v", refPath, err)
	}
	t, err := template.New(refPath).Parse(string(sample))
	if err != nil {
		return fmt.Errorf("Parsing %q: %v", refPath, err)
	}
	if _, err = parent.AddParseTree(refPath, t.Tree); err != nil {
		return fmt.Errorf("Adding parse tree of %q: %v", refPath, err)
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

func listMatchingFiles(dir string, ext string) (map[string]time.Time, error) {
	files := make(map[string]time.Time)
	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files[path] = fi.ModTime()
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Walking dir %s: %v", dir, err)
	}
	return files, err
}
