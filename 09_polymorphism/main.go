package main

import "fmt"

type Animal interface {
	Speak()
}

type dog struct {
	Name string
}

func (d *dog) Speak() {
	fmt.Printf("[Dog] %s says: Woof woof!\n", d.Name)
}

func NewDog(name string) Animal {
	return &dog{
		Name: name,
	}
}

type cat struct {
	Name string
}

func (c *cat) Speak() {
	fmt.Printf("[Cat] %s says: Miau miau!\n", c.Name)
}

func NewCat(name string) Animal {
	return &cat{
		Name: name,
	}
}

func main() {
	var animals [2]Animal

	animals[0] = NewDog("Sparky")
	animals[1] = NewCat("Andromeda")

	for _, a := range animals {
		a.Speak()
	}
}
