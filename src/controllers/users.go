package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyResquest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal(bodyResquest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	repository := repositories.NewUsersRepository(db)
	repository.Create(user)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding Users"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding User"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}
