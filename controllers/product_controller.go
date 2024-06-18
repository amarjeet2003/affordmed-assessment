package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"top-products-microservice/services"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["categoryname"]

	n, _ := strconv.Atoi(r.URL.Query().Get("top"))
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("minPrice"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("maxPrice"), 64)

	products, err := services.FetchTopProducts(category, n, minPrice, maxPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["productid"]

	product, err := services.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
