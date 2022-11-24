package routes

import (
  "github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
  UserRoutes(r)
  ProductRoutes(r) // Add this code
  ToppingRoutes(r)
  AuthRoutes(r)
}
