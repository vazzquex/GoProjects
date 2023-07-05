// main.go
package main

import (
	"net/http"

	"go-frontend/routes"
	templatehandler "go-frontend/templates"

	"github.com/gorilla/mux"
)

func main() {
	// Preprocesa las plantillas
	templatehandler.LoadTemplates("templates/*.html")

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	// r.HandleFunc("/products", routes.ProductsHandler)
	// r.HandleFunc("/articles", routes.ArticlesHandler)

	// Sirve los archivos estáticos
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
