package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/KristianXi3/crud/entity1"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandlerInterface interface {
	LoginsHandler(w http.ResponseWriter, r *http.Request)
}

type LoginHandler struct {
}

func NewLoginHandler() LoginHandlerInterface {
	return &LoginHandler{}
}

func (l *LoginHandler) LoginsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		LoginUser(w, r, id)
	}
}
func LoginUser(w http.ResponseWriter, r *http.Request, userid string) {
	var user entity1.User
	var cred entity1.Credentials
	ctx := context.Background()
	var jwtKey = []byte("my_secret_key")

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := SqlConnect.LoginsUser(ctx, userid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeJsonResp(w, statusSuccess, rows)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &entity1.Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}
