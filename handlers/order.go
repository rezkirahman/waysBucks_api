package handlers

import (
	"encoding/json"
	"net/http"

	//"strconv"
	dto "waysbucks/dto/result"
	//ordersdto "waysbucks/dto/orders"
	//"waysbucks/models"
	"waysbucks/repositories"
	//"github.com/go-playground/validator/v10"
	//"github.com/gorilla/mux"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) FindOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: orders}
	json.NewEncoder(w).Encode(response)
}
