package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KristianXi3/crud/entity1"
	"github.com/gorilla/mux"
)

type UserHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
}

func NewUserHandler() UserHandlerInterface {
	return &UserHandler{}
}
func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		// fmt.Println("TestGet")
		if id != "" { // get by id
			getUsersByIDHandler(w, r, id)
		} else { // get all
			h.getUsersHandler(w, r)
		}
	case http.MethodPost:
		createUsersHandler(w, r)
	case http.MethodPut:
		updateUserHandler(w, r, id)
	case http.MethodDelete:
		deleteUserHandler(w, r, id)
	}
}

func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	users, err := SqlConnect.GetUsers(ctx)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}
	writeJsonResp(w, statusSuccess, users)
}

func getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	if idInt, err := strconv.Atoi(id); err == nil {
		ctx := context.Background()
		users, err := SqlConnect.GetUserByID(ctx, idInt)
		if err != nil {
			writeJsonResp(w, statusError, err.Error())
			return
		}
		writeJsonResp(w, statusSuccess, users)
	}
}
func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var user entity1.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	users, err := SqlConnect.CreateUser(ctx, user)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}
	writeJsonResp(w, statusSuccess, users)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()

	if id != "" { // get by id
		decoder := json.NewDecoder(r.Body)
		var user entity1.User
		if err := decoder.Decode(&user); err != nil {
			w.Write([]byte("error decoding json body"))
			return
		}

		if idInt, err := strconv.Atoi(id); err == nil {
			if idInt != user.Id {
				writeJsonResp(w, statusError, "No ID not same")
				return
			} else if users, err := SqlConnect.GetUserByID(ctx, idInt); err != nil {
				writeJsonResp(w, statusError, err.Error())
				return
			} else if users.Id == 0 {
				writeJsonResp(w, statusError, "Data not exists")
				return
			} else {
				users, err := SqlConnect.UpdateUser(ctx, idInt, user)
				if err != nil {
					writeJsonResp(w, statusError, err.Error())
					return
				}
				writeJsonResp(w, statusSuccess, users)
			}
		}
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" { // get by id
		if idInt, err := strconv.Atoi(id); err == nil {
			users, err := SqlConnect.DeleteUser(ctx, idInt)
			if err != nil {
				writeJsonResp(w, statusError, err.Error())
				return
			}
			writeJsonResp(w, statusSuccess, users)
		}
	}
}
