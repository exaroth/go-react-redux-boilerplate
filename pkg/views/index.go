package views

import (
	"net/http"
)

var IndexView http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	templateData := map[string]interface{}{
		"test": 1,
	}
	err := RenderTemplate(w, "index", templateData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func RenderTemplate(w http.ResponseWriter, tplName string, tplData map[string]interface{}) error {
	// tpl := Config.GetTemplate(tplName)
	// if tpl == nil {
	// 	return fmt.Errorf("Template %s not found", tplName)
	// }
	// err := tpl.Execute(w, ctx)
	return nil
}
