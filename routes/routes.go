package routes

import (
	"top-products-microservice/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/categories/{categoryname}/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/categories/{categoryname}/products/{productid}", controllers.GetProductByID).Methods("GET")

	return router
}
