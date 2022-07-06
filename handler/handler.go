package handler

import (
	"CRUD/entity1"

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

func NewUserHandler() UserHandlerInterface {
	//return &UserHandler{postgrespool: postgrespool}
	return &users{}
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

// func UsersHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	switch r.Method {
// 	case http.MethodGet:
// 		if id != "" { // get by id
// 			getUsersByIDHandler(w, r, id)
// 		} else { // get all
// 			getUsersHandler(w, r)
// 		}
// 	case http.MethodPost:
// 		createUsersHandler(w, r)
// 	case http.MethodPut:
// 		updateUserHandler(w, r, id)
// 	case http.MethodDelete:
// 		deleteUserHandler(w, r, id)
// 	}
// }

// func getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	if v, ok := users[id]; ok {
// 		w.Header().Add("Content-Type", "application/json")
// 		json, _ := json.Marshal(v)
// 		w.Write(json)
// 	}
// }

// func (h *UsersHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	rows, err := h.postgrespool.Query(ctx, "select * from public.user")
// 	if err != nil {
// 		fmt.Println("query row error", err)
// 	}
// 	defer rows.Close()

// 	users := []*entity1.User{}
// 	for rows.Next() {
// 		var user entity1.User
// 		if serr := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt); serr != nil {
// 			fmt.Println("Scan error", serr)
// 		}
// 		users = append(users, &user)
// 	}

// 	jsonData, _ := json.Marshal(&users)
// 	w.Header().Add("Content-Type", "application/json")
// 	w.Write(jsonData)
// }
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, _ := json.Marshal(users)
	w.Write(json)
}

func getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
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
