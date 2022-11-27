package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/orders", h.FindOrders).Methods("GET")
	r.HandleFunc("/order/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/order", h.CreateOrder).Methods("POST")
	r.HandleFunc("/order/{}", h.UpdateOrder).Methods("PATCH")
	r.HandleFunc("/order/{id}", h.DeleteOrder).Methods("DELETE")
}
