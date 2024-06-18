package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"top-products-microservice/models"
	"top-products-microservice/utils"
)

var exampleProducts = []models.Product{
	{
		ID:           "1",
		ProductName:  "Example Product 1",
		Price:        99.99,
		Rating:       4.5,
		Discount:     10,
		Availability: "yes",
		Company:      "AMZ",
	},
	{
		ID:           "2",
		ProductName:  "Example Product 2",
		Price:        149.99,
		Rating:       4.8,
		Discount:     20,
		Availability: "yes",
		Company:      "FLP",
	},
}

func FetchTopProducts(category string, n int, minPrice, maxPrice float64) ([]models.Product, error) {
	companies := []string{"AMZ", "FLP", "SNP", "MYN", "AZO"}
	var allProducts []models.Product

	for _, company := range companies {
		url := fmt.Sprintf("%s%s/categories/%s/products?top=%d&minPrice=%.2f&maxPrice=%.2f", utils.BaseURL, company, category, n, minPrice, maxPrice)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var products []models.Product
		if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
			return nil, err
		}

		for i := range products {
			products[i].ID = models.GenerateProductID(company, category, products[i].ProductName)
			products[i].Company = company
		}

		allProducts = append(allProducts, products...)
	}

	if len(allProducts) == 0 {
		return exampleProducts, nil
	}

	return allProducts, nil
}

func GetProductByID(productID string) (models.Product, error) {
	for _, p := range exampleProducts {
		if p.ID == productID {
			return p, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}
