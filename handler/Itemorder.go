package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
func (h *UserHandler) OrdersHandler(w http.ResponseWriter, r *http.Request) {
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
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var order entity1.Order
	var item entity1.Items
	if err := decoder.Decode(&order); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	order, err := SqlConnect.CreateOrder(ctx, order, item)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}
	writeJsonResp(w, statusSuccess, order)
}

func updateOrdersHandler(w http.ResponseWriter, r *http.Request, orderid string) {
	ctx := context.Background()

	if orderid != "" { // get by id
		decoder := json.NewDecoder(r.Body)
		var Order entity1.Order
		var item entity1.Items
		if err := decoder.Decode(&Order); err != nil {
			w.Write([]byte("error decoding json body"))
			return
		}

		if idInt, err := strconv.Atoi(orderid); err == nil {
			if idInt != Order.Order_id {
				writeJsonResp(w, statusError, "No ID not same")
				return
			} else if order, err := SqlConnect.GetOrderByID(ctx, idInt); err != nil {
				writeJsonResp(w, statusError, err.Error())
				return
			} else if order.Order_id == 0 {
				writeJsonResp(w, statusError, "Data not exists")
				return
			} else {
				order, err := SqlConnect.UpdateOrder(ctx, idInt, Order, item)
				if err != nil {
					writeJsonResp(w, statusError, err.Error())
					return
				}
				writeJsonResp(w, statusSuccess, order)
			}
		}
	}
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
