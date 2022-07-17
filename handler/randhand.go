package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KristianXi3/crud/entity1"
)

type RandUserHandInterface interface {
	RandUserHandler(w http.ResponseWriter, r *http.Request)
}

type RandUserHandler struct {
}

func RandUserHandlerFunc() RandUserHandInterface {
	return &RandUserHandler{}
}

func (u *RandUserHandler) RandUserHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	// id := params["id"]
	switch r.Method {
	case http.MethodGet:
		// fmt.Println("TestGet")
		// if id != "" { // get by id
		// 	getOrdersByIDHandler(w, r, id)
		// } else { // get all
		getRandData(w, r)
	}
	//getRandData(w, r)

}

func getRandData(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var users []entity1.RandGenUser
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal(body, &users)

	tpl, err := template.ParseFiles("html/template.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, users)

}
