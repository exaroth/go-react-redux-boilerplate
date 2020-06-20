package views

import (
	"fmt"
	"io"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
)

// RenderTemplate will write given template with given context into the response.
func RenderTemplate(w io.Writer, tplN string, tplD map[string]interface{}) error {
	tpl := config.Config.GetTemplate(tplN)
	if tpl == nil {
		return fmt.Errorf("Template %s not found", tplN)
	}
	return tpl.Execute(w, tplD)
}
