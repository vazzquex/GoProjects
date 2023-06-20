package routes

import (
	"encoding/json"
	"net/http"

	"github.com/vazzquex/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user = models.User

	json.NewDecoder(r.Body).Decode(&user)

	w.Write([]byte("Post"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
