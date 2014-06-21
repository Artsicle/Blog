package helpers

import (
	"fmt"
	"html/template"
)

type Helper struct{}

type Linker interface {
	LinkTitle() string
	Url() string
}

func (h Helper) LinkTo(a Linker) template.HTML {
	content := fmt.Sprintf("<a href='%s'>%s</a>", a.Url(), a.LinkTitle())
	return template.HTML(content)
}
