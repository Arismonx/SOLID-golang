package main

import (
	"encoding/json"
	"fmt"
)

type (
	Product struct {
		ID              uint    `json:id`
		Name            string  `json:name`
		Price           float32 `json:price`
		Category        string  `json:category`
		DiscountedPrice float32 `json:discountedPrice`
	}

	ProductService struct {
		saves []Product
	}
	Discountservice struct{}
	ReportService   struct{}
)

func (pro *ProductService) SaveProduct(product []Product) {
	pro.saves = append(pro.saves, product...)
	fmt.Printf("Product Save: %v", product)

}

func (dis Discountservice) Discount(product []Product) {
	var s float32
	for i := range product {
		if product[i].Category == "Electronics" {
			s = 10.0 / 100.0 * product[i].Price
			s = product[i].Price - s
			product[i].DiscountedPrice = s
		}
		if product[i].Category == "Clothing" {
			s = 5.0 / 100.0 * product[i].Price
			s = product[i].Price - s
			product[i].DiscountedPrice = s

		}
		if product[i].Category == "Books" {
			s = 2.0 / 100.0 * product[i].Price
			s = product[i].Price - s
			product[i].DiscountedPrice = s
		}
	}
}

func (r ReportService) Report(product []Product) string {
	jsondata, _ := json.MarshalIndent(product, "", " ")
	return string(jsondata)
}

func main() {
	product := []Product{
		{ID: 1, Name: "Laptop", Price: 1000, Category: "Electronics"},
		{ID: 2, Name: "Shirt", Price: 50, Category: "Clothing"},
		{ID: 3, Name: "Book", Price: 30, Category: "Books"},
	}

	productService := ProductService{}
	productService.SaveProduct(product)

	dis := &Discountservice{}
	dis.Discount(product)

	report := ReportService{}
	fmt.Println(report.Report(product))

}
