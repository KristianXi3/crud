package main

import (
	"CRUD/entity1"
	"CRUD/handler"
	"CRUD/service"
	"assignment/crud/handler"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":8080"
var server = "localhost"
var port = 1433
var database = "GoLang"

var db *sql.DB

func main() {
	connString := fmt.Sprintf("server=%s;port=%d; database = %s",
		server, port, database, "trusted_connection=yes")
	sql := database.ConnectSQL(connString)
	handler.SqlConnect = sql
	r := mux.NewRouter()
	userHandler := handler.NewUserHandler()
	//r.HandleFunc("/", greet)
	r.HandleFunc("/register", userRegister).MethodPost
	r.HandleFunc("/user", userHandler.UsersHandler)
	r.HandleFunc("/user/{id}", userHandler.UsersHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func userRegister(w http.ResponseWriter, r *http.Request) {
	userSvc := service.NewUserService()

	decoder := json.NewDecoder(r.Body)
	var newUser entity1.User
	if err := decoder.Decode(&newUser); err != nil {
		w.WriteHeader(201)
		w.Write([]byte("error decoding json body"))
		return
	}

	if user, err := userSvc.Register(&newUser); err != nil {
		fmt.Printf("Error when register user: %+v \n", err)
		w.WriteHeader(201)
		w.Write([]byte("Error when register user"))
		return
	} else {
		m, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("Error when register user: %+v \n", err)
			w.WriteHeader(201)
			w.Write([]byte("Error when register user"))
		}

		fmt.Printf("Success register user: %+v \n", user)
		fmt.Println("----------------------------------")
		w.Header().Add("Content-Type", "application/json")
		w.Write(m)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var user entity1.User
		if err := decoder.Decode(&user); err != nil {
			w.Write([]byte("error decoding json body"))
			return
		}

		userSvc := service.NewUserService()
		res, err := userSvc.register(&user)

		jData, _ := json.Marshal(res)

		w.Header().Add("Content-Type", "application/json")
		w.Write(jData)

		if err != nil {
			w.Write([]byte("error decoding json body"))

		}

	}

}

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
