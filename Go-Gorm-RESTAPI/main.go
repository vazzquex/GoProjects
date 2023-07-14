package main

import (
	"net/http"

	"go-gorm-restapi/db"
	"go-gorm-restapi/models"
	"go-gorm-restapi/routes"

	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	//Users routes

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//tasks routes
	r.HandleFunc("/tasks", routes.GetTasksHadnle).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHadnle).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHAndler).Methods("DELETE")

	// Manejador para servir archivos estáticos
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler)

	http.ListenAndServe(":3000", r)

}
