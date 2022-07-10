package handler

import (
	"github.com/KristianXi3/crud/entity1"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = map[int]entity1.User{
	1: {
		Id:       1,
		Username: "andi123",
		Email:    "andi123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	2: {
		Id:       2,
		Username: "budi123",
		Email:    "budi123@gmail.com",
		Password: "password123",
		Age:      9,
	},
	3: {
		Id:       3,
		Username: "cantya123",
		Email:    "cantya123@gmail.com",
		Password: "password123",
		Age:      9,
	},
}

type UserHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	//postgrespool *pgxpool.Pool
}

func NewUserHandler() UserHandlerInterface {
	//return &UserHandler{postgrespool: postgrespool}
	return &UserHandler{}
}
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	switch r.Method {
	case http.MethodGet:
		if id != "" { // get by id
			getUsersByIDHandler(w, r, id)
		} else { // get all
			getUsersHandler(w, r)
		}
	case http.MethodPost:
		createUsersHandler(w, r)
	case http.MethodPut:
		updateUserHandler(w, r, id)
	case http.MethodDelete:
		deleteUserHandler(w, r, id)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, _ := json.Marshal(users)
	w.Write(json)
}

func getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	users, err := SqlConnect.Getuser
	if idInt, err := strconv.Atoi(id); err == nil {
		if user, ok := users[idInt]; ok {
			jsonData, _ := json.Marshal(user)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
			return
		} else {
			w.Write([]byte("No user found for given id"))
			return
		}
	}
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity1.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	if _, found := users[user.Id]; found {
		w.Write([]byte("User with given id already exists"))
		return
	}

	users[user.Id] = user
	var usersSlice []entity1.User
	for _, v := range users {
		usersSlice = append(usersSlice, v)
	}
	jsonData, _ := json.Marshal(&usersSlice)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		if idInt, err := strconv.Atoi(id); err == nil {
			decoder := json.NewDecoder(r.Body)
			var user entity1.User
			if err := decoder.Decode(&user); err != nil {
				w.Write([]byte("error decoding json body"))
				return
			}

			users[idInt] = user
			jsonData, _ := json.Marshal(&user)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		if idInt, err := strconv.Atoi(id); err == nil {
			delete(users, idInt)
			w.Write([]byte(fmt.Sprintf("User %d deleted", idInt)))
		}
	}
}
