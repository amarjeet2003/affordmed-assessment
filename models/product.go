package models

import "fmt"

type Product struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"productName"`
	Price        float64 `json:"price"`
	Rating       float64 `json:"rating"`
	Discount     int     `json:"discount"`
	Availability string  `json:"availability"`
	Company      string  `json:"company"`
}

func GenerateProductID(company, category, productName string) string {
	// Generate custom ID
	return fmt.Sprintf("%s-%s-%s", company, category, productName)
}
