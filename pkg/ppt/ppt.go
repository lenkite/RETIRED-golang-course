package ppt

import (
	"fmt"
	"html/template"
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

	// refPaths are the child template paths embedded in the presentation
	refPaths []string

	// root is the html template of the presentation
	root *template.Template

	// refs are the html templates of the references
	refs []*template.Template
}

// New returns a new Presentation for the given configuration
// All templates are loaded and the presentation can be written to an io.Writer
func New(config Config) (*Presentation, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("Loading ppt from: %q: %v", path, err)
	}
	return
}
