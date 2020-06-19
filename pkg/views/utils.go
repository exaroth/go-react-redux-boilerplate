package views

import (
	"fmt"
	"net/http"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
)

// RenderTemplate will write given template with given context into the response.
func RenderTemplate(w http.ResponseWriter, tplN string, tplD map[string]interface{}) error {
	tpl := config.Config.GetTemplate(tplN)
	if tpl == nil {
		return fmt.Errorf("Template %s not found", tplN)
	}
	return tpl.Execute(w, tplD)
}
