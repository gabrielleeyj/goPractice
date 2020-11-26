package main

import "fmt"

type Engine struct {
	Trademark  string
	HorsePower float64
}

type Seat struct {
	Trademark string
}

func NewEngine(tm string, hp float64) Engine {
	return Engine{
		Trademark:  tm,
		HorsePower: hp,
	}
}

// Composition over inheritance

type Car struct {
	engine Engine
	seat   *Seat // nil
}

func NewCar(engine Engine) Car {
	return Car{
		engine: engine,
	}
}

func main() {
	e := NewEngine("Audi", 1200.0)
	c := NewCar(e)

	fmt.Printf("Car: %+v\n", c)
}
