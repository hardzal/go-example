package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hardzal/go-example/go-books/pkg/models"
	"github.com/hardzal/go-example/go-books/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var NewUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing user json")
	}

	userDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	User := &models.User{}
	utils.ParseBody(r, User)
	password, _ := bcrypt.GenerateFromPassword([]byte(User.Password), 14)
	User.Password = string(password)
	u := User.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var User = &models.User{}
	utils.ParseBody(r, User)
	vars := mux.Vars(r)
	userId := vars["userId"]

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	userDetails, db := models.GetUserById(ID)
	if User.Username != "" {
		userDetails.Username = User.Username
	}

	if User.Name != "" {
		userDetails.Name = User.Name
	}

	if User.Email != "" {
		userDetails.Email = User.Email
	}

	if User.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(User.Password), 14)
		if err == nil {
			userDetails.Password = string(password)
		}
	}

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
