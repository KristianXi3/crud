package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/KristianXi3/crud/entity1"

	"github.com/gorilla/mux"
)

type OrderHandlerInterface interface {
	OrdersHandler(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
}

func NewOrderHandler() OrderHandlerInterface {
	return &OrderHandler{}
}
func (h *OrderHandler) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		// fmt.Println("TestGet")
		if id != "" { // get by id
			getOrdersByIDHandler(w, r, id)
		} else { // get all
			h.getOrdersHandler(w, r)
		}
	case http.MethodPost:
		createOrdersHandler(w, r)
	case http.MethodPut:
		updateOrdersHandler(w, r, id)
	case http.MethodDelete:
		deleteOrdersHandler(w, r, id)
	}
}

func (h *OrderHandler) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	order, err := SqlConnect.GetOrders(ctx)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}
	writeJsonResp(w, statusSuccess, order)
}

func getOrdersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	if idInt, err := strconv.Atoi(id); err == nil {
		ctx := context.Background()
		order, err := SqlConnect.GetOrderByID(ctx, idInt)
		if err != nil {
			writeJsonResp(w, statusError, err.Error())
			return
		}
		writeJsonResp(w, statusSuccess, order)
	}
}
func createOrdersHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var order entity1.Order

	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelfunc()
	result, err := SqlConnect.CreateOrder(ctx, order)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result))
}

func updateOrdersHandler(w http.ResponseWriter, r *http.Request, orderid string) {
	if orderid != "" {
		if idInt, err := strconv.Atoi(orderid); err == nil {
			decoder := json.NewDecoder(r.Body)
			var orderSlice entity1.Order2
			if err := decoder.Decode(&orderSlice); err != nil {
				w.Write([]byte("Error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := SqlConnect.UpdateOrder(ctx, idInt, orderSlice)

			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(result))
			return
		}
	}
	w.Write([]byte("Invalid parameter"))
	// return
}

func deleteOrdersHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" { // get by id
		if idInt, err := strconv.Atoi(id); err == nil {
			orders, err := SqlConnect.DeleteOrder(ctx, idInt)
			if err != nil {
				writeJsonResp(w, statusError, err.Error())
				return
			}
			writeJsonResp(w, statusSuccess, orders)
		}
	}
}
