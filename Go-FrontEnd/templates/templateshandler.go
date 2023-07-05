// templatehandler/templatehandler.go
package templatehandler

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	once      sync.Once
	templates *template.Template
)

func LoadTemplates(pattern string) {
	once.Do(func() {
		templates = template.Must(template.ParseGlob(pattern))
	})
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return templates.ExecuteTemplate(w, tmpl, data)
}
