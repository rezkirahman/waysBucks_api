package routes

import (
  "waysbucks/handlers"
  "waysbucks/pkg/middleware"
  "waysbucks/pkg/mysql"
  "waysbucks/repositories"

  "github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
  productRepository := repositories.RepositoryProduct(mysql.DB)
  h := handlers.HandlerProduct(productRepository)

  r.HandleFunc("/products", middleware.Auth(h.FindProducts)).Methods("GET") // add this code
  r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
  r.HandleFunc("/product", middleware.Auth(h.CreateProduct)).Methods("POST") // add this code
}