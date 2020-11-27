package main

import "fmt"

func main() {
	var list []Product

	fmt.Println("List t=0:", list)

	list = AddProduct(AddProductInput{
		List:     list,
		Name:     "Pepito",
		Price:    1200,
		Discount: 2,
	})

	fmt.Println("List t=1:", list)
}

// AddProductInput is the input DTO of the AddProduct function.
type AddProductInput struct {
	List     []Product
	Name     string
	Price    float64
	Discount float64
	Value    float64
}

// AddProductOutput is the output DTO of the AddProduct function.
type AddProductOutput []Product

// Product Structure contains abstraction of the product information.
type Product struct {
	Name     string
	Price    float64
	Discount float64
}

// AddProduct function takes the input AddProductInput which is the DTO of the AddProduct function.
func AddProduct(input AddProductInput) []Product {
	var out []Product

	out = append(input.List, Product{
		Name:     input.Name,
		Price:    input.Price,
		Discount: input.Discount,
	})

	return out
}
