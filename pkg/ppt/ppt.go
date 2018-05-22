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
	//refPaths map[string]time.Time
	refs map[string]Ref
}

// Ref encapsulates a HTML template reference, its path and last modified time.
type Ref struct {
	Path    string
	ModTime time.Time
	tmpl    *template.Template
}

// New returns a new Presentation for the given configuration
// All templates are loaded and the presentation can be written to an io.Writer
func New(config Config) (*Presentation, error) {
	var ppt Presentation
	ppt.Config = config
	if !strings.HasPrefix(ppt.RefExt, ".") {
		fmt.Println("Adding dot to RefExt")
		ppt.RefExt = "." + ppt.RefExt
	}
	if err := ppt.createOrUpdate(); err != nil {
		return nil, err
	}
	return &ppt, nil
}

func (ppt *Presentation) createOrUpdate() error {
	if ppt.refs == nil {
		ppt.refs = make(map[string]Ref)
	}
	if err := addUpdateRefs(ppt.refs, ppt.RefDir, ppt.RefExt); err != nil {
		return fmt.Errorf("Loading templates from: %q: %v", ppt.RefDir, err)
	}
	return ppt.createRootTemplate()
}

//WriteTo Writes the serialized presentatio html to given writer
func (ppt *Presentation) WriteTo(w io.Writer) (n int64, err error) {
	if err := ppt.ReloadChanges(); err != nil {
		return 0, err
	}
	var buf bytes.Buffer
	if err := ppt.root.Execute(&buf, ppt.Name); err != nil {
		return 0, fmt.Errorf("Writing %q to %v: %v", ppt.Name, w, err)
	}
	return buf.WriteTo(w)
}

func (ppt *Presentation) ReloadChanges() error {
	if !ppt.HotReload {
		return nil
	}
	return ppt.createOrUpdate()
}

func (ppt *Presentation) createRootTemplate() error {
	fmt.Printf("parsing files from %q\n", ppt.Path)
	tmpl, err := template.ParseFiles(ppt.Path)
	fmt.Printf("Parsed template %v\n", tmpl)
	if err != nil {
		return fmt.Errorf("Loading ppt from: %q: %v", ppt.Path, err)
	}
	for _, ref := range ppt.refs {
		if _, err = tmpl.AddParseTree(ref.Path, ref.tmpl.Tree); err != nil {
			return fmt.Errorf("Added parse tree of %q: %v", ref.Path, err)
		}
	}
	ppt.root = tmpl
	return nil
}

func addUpdateRefs(refs map[string]Ref, dir string, ext string) error {
	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if filepath.Ext(path) != ext {
			return nil
		}
		pathModTime := fi.ModTime()
		if r, ok := refs[path]; ok { //ref already in map
			if r.ModTime.Equal(pathModTime) {
				return nil //if ref time hasn't change return
			}
		}
		fmt.Printf("Loading Ref %q...", path)
		refTmpl, err := createRefTemplate(path, fi.ModTime())
		if err != nil {
			return err
		}
		refs[path] = Ref{Path: path, ModTime: pathModTime, tmpl: refTmpl}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Walking dir %q: %v", dir, err)
	}
	return nil
}

func createRefTemplate(refPath string, modTime time.Time) (*template.Template, error) {
	sample, err := ioutil.ReadFile(refPath)
	if err != nil {
		return nil, fmt.Errorf("Cannot read go file %q: %v", refPath, err)
	}
	t, err := template.New(refPath).Parse(string(sample))
	if err != nil {
		return nil, fmt.Errorf("Parsing %q: %v", refPath, err)
	}
	return t, nil
}

// func addRefTemplates(parent *template.Template, dir string, ext string) (map[string]time.Time, error) {
// 	refPaths := make(map[string]time.Time)
// 	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
// 		if filepath.Ext(path) != ext {
// 			return nil
// 		}
// 		refPaths[path] = fi.ModTime()
// 		return addRefTemplate(parent, path)
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("Walking dir %q: %v", dir, err)
// 	}
// 	return refPaths, nil
// }

// func addRefTemplate(parent *template.Template, refPath string) error {
// 	sample, err := ioutil.ReadFile(refPath)
// 	if err != nil {
// 		return fmt.Errorf("Cannot read go file %q: %v", refPath, err)
// 	}
// 	t, err := template.New(refPath).Parse(string(sample))
// 	if err != nil {
// 		return fmt.Errorf("Parsing %q: %v", refPath, err)
// 	}
// 	if _, err = parent.AddParseTree(refPath, t.Tree); err != nil {
// 		return fmt.Errorf("Adding parse tree of %q: %v", refPath, err)
// 	}
// 	return nil
// }

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
