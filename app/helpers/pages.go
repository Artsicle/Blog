package helpers

import (
	"fmt"
	"html/template"
)

type Helper struct{}

func (h Helper) LinkTo(name, url string) template.HTML {
	var content string = fmt.Sprintf("<a href='%s'>%s</a>", url, name)
	return template.HTML(content)
}
