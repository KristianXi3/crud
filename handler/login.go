package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KristianXi3/crud/entity1"
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
	// params := mux.Vars(r)
	// id := params["id"]

	switch r.Method {
	case http.MethodPost:
		LoginUser(w, r)
	}
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var cred entity1.Credentials
	ctx := context.Background()

	err := json.NewDecoder(r.Body).Decode(&cred)
	fmt.Println(cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := SqlConnect.LoginsUser(ctx, cred)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeJsonResp(w, statusSuccess, rows)

}
