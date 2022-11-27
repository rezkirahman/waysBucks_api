package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	orderdto "waysbucks/dto/order"
	dto "waysbucks/dto/result"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
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

func (h *handlerOrder) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseOrder(order)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := new(orderdto.CreateOrderRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	//request product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.OrderRepository.ProductOrder(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Product not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//request topping id
	toppings, err := h.OrderRepository.ToppingOrder(request.ToppingID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Topping not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var totalTopping = 0
	for _, i := range toppings {
		totalTopping += i.Price
	}
	var subTotal = request.Qty * (product.Price + totalTopping)

	order := models.Order{
		UserID:    userId,
		Qty:       request.Qty,
		ProductID: product.ID,
		Topping:   toppings,
		Subtotal:  subTotal,
	}

	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *handlerOrder) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	cart, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.OrderRepository.DeleteOrder(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseOrder(u models.Order) models.Order {
	return models.Order{
		ID:       u.ID,
		Qty:      u.Qty,
		Subtotal: u.Subtotal,
		Product:  u.Product,
		Topping:  u.Topping,
	}
}
