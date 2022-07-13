package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KristianXi3/crud/DB"
	"github.com/KristianXi3/crud/handler"

	"github.com/gorilla/mux"
)

var PORT = ":8080"
var server = "localhost"
var port = 1433
var database = "GoLang"

func main() {
	connString := fmt.Sprintf("server=%s;port=%d; database=%s ;trusted_connection=yes",
		server, port, database)
	sql := DB.ConnectSQL(connString)
	handler.SqlConnect = sql
	r := mux.NewRouter()
	userHandler := handler.NewUserHandler()
	//r.HandleFunc("/", greet)
	r.HandleFunc("/users", userHandler.UsersHandler)
	r.HandleFunc("/users/{id}", userHandler.UsersHandler)
	orderHandler := handler.NewOrderHandler()
	r.HandleFunc("/order", orderHandler.OrdersHandler)
	r.HandleFunc("/order/{id}", orderHandler.OrdersHandler)
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
