package main

import (
	"CRUD/entity1"
	"CRUD/service"
	"assignment/crud/handler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

func main() {
	r := mux.NewRouter()
	userHandler := handler.NewUserHandler()
	//r.HandleFunc("/", greet)
	r.HandleFunc("/register", userRegister).MethodPost
	r.HandleFunc("/user", userHandler.UsersHandler)
	r.HandleFunc("/user/{id}", userHandler.UsersHandler)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
	// r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET"
	// 	// an example API handler
	// 	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	// })

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	// userSvc := service.NewUserService()

	// userSvc.Register(&entity1.User{
	// 	Id:        1,
	// 	Username:  "Kristian",
	// 	Email:     "email@email.com",
	// 	Password:  "Password",
	// 	Age:       17,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })

}

func userRegister(w http.ResponseWriter, r *http.Request) {
	userSvc := service.NewUserService()
	// newUser := &entity.User{
	// 	Id:        1,
	// 	Username:  "david123",
	// 	Email:     "david123@gmail.com",
	// 	Password:  "Passdav!d",
	// 	Age:       17,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

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
		// userSvc1 := userSvc.Register(&user)
		// json, _ := json.Marshal(userSvc1)
		// w.Write(json)
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

// func UsersHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	checkHandler := handler.UserHandlerInterface
// 	switch r.Method {
// 	case http.MethodGet:
// 		if id != "" { // get by id
// 			checkHandler.getUsersByIDHandler(w, r, id)
// 		} else { // get all
// 			checkHandler.getcompleteuser(w, r)
// 		}
// 	case http.MethodPost:
// 		checkHandler.createUsersHandler(w, r)
// 	case http.MethodPut:
// 		checkHandler.updateUserHandler(w, r, id)
// 	case http.MethodDelete:
// 		checkHandler.deleteUserHandler(w, r, id)
// 	}
// }

// func getcompleteuser(w http.ResponseWriter, r *http.Request) {
// 	x := []entity1.User{}
// 	for _, val := range users {
// 		x = append(x, val)
// 	}
// 	check, _ := json.Marshal(x)
// 	w.Header().Add("Content-Type", "application/json")
// 	w.Write(check)
// }
// func createUsersHandler(w http.ResponseWriter, r *http.Request){

// }
