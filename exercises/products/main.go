package main

import "fmt"

// Product is an abstraction of ...
type Product struct {
	// Name is the name of the product.
	Name string

	// Trademark is the name of the company that created this product.
	Trademark string

	// Price is the price that the user will need to pay to acquire this product.
	Price float32

	// Discount is a percentage of discount that will be applied to this product.
	Discount float32
}

func (p *Product) ApplyDiscount(discount float32) {
	if discount < 0 {
		return
	}
	p.Discount = discount
}

func (p *Product) ChangePrice(price float32) {
	p.Price = price
}

func ChangeAnotherPrice() {

}

func main() {
	var products []Product

	product := NewProduct("Marcos", "Marcos LTD", 1.4)

	fmt.Println("List [t=0]:", products)

	products = append(products, product)
	fmt.Println("List [t=1]:", products)

	product = NewProduct("Marcos", "Marcos LTD", 1.5)
	product.ApplyDiscount(1.0)
	product.ChangePrice(2)

	products = append(products, product)
	fmt.Println("List [t=2]:", products)

}

// NewProduct initializes a new Product from the given name, trademark and price.
func NewProduct(name, tm string, price float32) Product {
	return Product{
		Name:      name,
		Trademark: tm,
		Price:     price,
		Discount:  0.0,
	}
}

// NewProductWithDiscount initializes a new product from the given anme, trademark and price, but also applies a discount.
func NewProductWithDiscount(name, tm string, price, discount float32) Product {
	p := NewProduct(name, tm, price)
	p.Discount = discount
	return p
}
