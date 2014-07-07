package models

import (
	"fmt"
	"github.com/revel/revel"
	"html/template"
)

type Linker interface {
	LinkTitle() string
	Url() string
}

// add custom helpers functions for models to revel's template helper list
func init() {
	revel.TemplateFuncs["linkto"] = func(a Linker) template.HTML {
		content := fmt.Sprintf("<a href='%s'>%s</a>", a.Url(), a.LinkTitle())
		return template.HTML(content)
	}
}
