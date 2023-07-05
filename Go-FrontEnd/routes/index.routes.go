// routes/views.router.go
package routes

import (
	"net/http"

	templatehandler "go-frontend/templates"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Página de inicio",
		Message: "¡Bienvenido a mi sitio web!",
	}

	err := templatehandler.RenderTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, "Ocurrió un error interno", http.StatusInternalServerError)
		return
	}
}
