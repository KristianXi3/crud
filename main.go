package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KristianXi3/crud/DB"
	"github.com/KristianXi3/crud/entity1"
	"github.com/KristianXi3/crud/handler"
	"github.com/KristianXi3/crud/service"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var PORT = ":8080"
var server = "localhost"
var port = 1433
var database = "GoLang"

const secretkey = "jwtsecretkey111"

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
	//enterHandler := handler.NewEnterHandler()
	r.HandleFunc("/users/register", userHandler.UsersHandler) //Isi register
	r.HandleFunc("/users/login", userHandler.UsersHandler)    //Isi login
	orderHandler := handler.NewOrderHandler()
	r.HandleFunc("/order", orderHandler.OrdersHandler)
	r.HandleFunc("/order/{id}", orderHandler.OrdersHandler)
	randHandler := handler.RandUserHandlerFunc()
	r.HandleFunc("/rand", randHandler.RandUserHandler)
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
func Login(w http.ResponseWriter, r *http.Request) {
	var authDetails entity1.Credentials
	err := json.NewDecoder(r.Body).Decode(&authDetails)

	service.NewUserService().Register()

	if err != nil {
		var err entity1.Error
		err = SetError(err, "Error in reading auth")
		json.NewEncoder(w).Encode(err)
		return
	}

	var user entity1.User

	query := "SELECT username, password from user where username = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	fmt.Printf(authDetails.Username)
	rows := handler.SqlConnect.SqlDb.QueryRowContext(ctx, query, authDetails.Username)
	if err != nil {
		panic(err)
	}
	rows.Scan(&user.Username, &user.Password)

	check := CheckPasswordHash(authDetails.Password, user.Password)
	if !check {
		var err entity1.Error
		err = SetError(err, "username or password is incorrect")
		json.NewEncoder(w).Encode(err)
		return
	}
	validToken, err := GenerateJWT(user.Username)
	if err != nil {
		var err entity1.Error
		err = SetError(err, "Failed to generate token")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token entity1.Token
	token.Username = user.Username
	token.TokenString = validToken
	json.NewEncoder(w).Encode(token)
	return
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func SetError(err entity1.Error, message string) entity1.Error {
	err.IsError = true
	err.Message = message
	return err
}
func GenerateJWT(username string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
