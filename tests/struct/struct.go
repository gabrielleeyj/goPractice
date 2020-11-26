// the idea is to store an array of structs
// first create a struct type
// create an array which will store all the struct variables
// print out a list from the array
// from the array print out via index the struct variables.

package main

import "fmt"

// Product is an abstraction of... -> domain - model
type Product struct {
	productName     string
	productPrice    float32
	productCategory string
}

// DTO method to create product
// type CreateProductInput struct {
// 	List           []Product
// 	Name, Category string
// 	Price          float32
// }

// type CreateProductOutput []Product

// func CreateProduct(input CreateProductInput) CreateProductOutput {
// 	var out CreateProductOutput

// 	out = append(input.List, Product{
// 		productName:     input.Name,
// 		productCategory: input.Category,
// 		productPrice:    input.Price,
// 	})
// 	return out
// }

// Products Global Variable
// if declared within main function it will only be usable in main()
var Products []Product

func prod() {
	// create a variable that will store as a array of products
	// var Products []Product

	// struct literal declaration
	item1 := Product{
		productName:     "Bread",
		productPrice:    1.25,
		productCategory: "Food",
	}
	fmt.Println(item1)

	// Create the first struct to initialize
	// then append it to the array.
	pl := newProduct("Fish", "Food", 20.0)
	Products = append(Products, pl)

	fmt.Println("T[001] ", Products)
	// Create 2nd struct and append it to the array.
	pl = newProduct("Water", "Drink", 0.5)
	Products = append(Products, pl)

	fmt.Println("T[002] ", Products)
	// Create 3rd struct and append it to the array.
	pl = newProduct("Detergent", "Household", 10.0)
	Products = append(Products, pl)

	fmt.Println("T[003] ", Products)

	// if we never initialise the product it will not be a valid input
	newProduct("Soap", "Bath", 0.5)

	fmt.Println("T[004]  No initialization to new items")
	fmt.Println("T[004] ", Products)

	// if we initialise but not append
	newProduct("Water", "Drink", 0.50)

	fmt.Println("T[005]  No append will not return result in array")
	fmt.Println("T[005] ", Products)

	// using addProduct function
	addProduct("Coke", "Drink", 2.50)
	addProduct("Soap", "Bath", 0.5)
	fmt.Println("Add Function Test", Products)

	// Loop over all indexes in the slice.
	// ... Print all struct data.
	listProduct()

	sortProduct("Food")

	sortProduct("Drink")

	// 	for i := range Products {
	// 		Products := Products[i]
	// 		fmt.Printf("\nName: %v \nPrice: $ %v \nCategory: %v\n===========\n", Products.productName, Products.productPrice, Products.productCategory)
	// 	}
}

// function to create new product and appends to array Products
// addProduct takes in 3 parameters
// returns the parameters to Product
func addProduct(name, category string, price float32) Product {
	pl := Product{
		productName:     name,
		productCategory: category,
		productPrice:    price,
	}
	Products = append(Products, pl)
	return pl

}

// old function
func newProduct(name, category string, price float32) Product {
	return Product{
		productName:     name,
		productCategory: category,
		productPrice:    price,
	}
}

// function to show list of items
// listProduct takes in an array of Product Struct
// and prints out the Products in the following format:
//
// Name:
// Price:
// Category:
// ================

func listProduct() {
	for i := range Products {
		Products := Products[i]
		fmt.Printf("\nName: %v \nPrice: $ %v \nCategory: %v\n===========\n", Products.productName, Products.productPrice, Products.productCategory)
	}
}

// function to sort items
// sortProduct takes in a string parameter
// and iterates through the array of product to print out the product based on category
func sortProduct(category string) {
	fmt.Println("Sorting by Type: ", category)
	for i := range Products {
		Products := Products[i]
		if category == Products.productCategory {
			fmt.Printf("\nName: %v \nPrice: $ %v \nCategory: %v\n", Products.productName, Products.productPrice, Products.productCategory)
		}
	}
	fmt.Println("\n===========")
}
