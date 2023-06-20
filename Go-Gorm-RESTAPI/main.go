package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vazzquex/go-gorm-restapi/db"
	"github.com/vazzquex/go-gorm-restapi/routes"
)

func main() {

	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)

}
