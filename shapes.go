package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
	cost   float32
}

func (c *Circle) CulcArea() float32 {
	return c.radius * c.radius * math.Pi
}

func (c *Circle) CostPerCm2() float32 {
	return c.cost / c.CulcArea()
}

type Rectangle struct {
	sideA float32
	sideB float32
}

func (r *Rectangle) CulcArea() float32 {
	return r.sideA * r.sideB
}

type AreaCulcable interface {
	CulcArea() float32
}

func CulcShapeArea(i AreaCulcable) float32 {
	return i.CulcArea()
}

func main() {
	pizzaSmall := &Circle{radius: 11, cost: 359}
	pizzaMiddle := &Circle{radius: 15, cost: 559}
	pizzaLarge := &Circle{radius: 20, cost: 829}
	rectangle := &Rectangle{sideA: 5, sideB: 7}

	fmt.Println("Small pizza: area ", CulcShapeArea(pizzaSmall), ",cost per cm2 ", pizzaSmall.CostPerCm2())
	fmt.Println("Medium pizza: area ", CulcShapeArea(pizzaMiddle), ",cost per cm2 ", pizzaMiddle.CostPerCm2())
	fmt.Println("Large pizza area", CulcShapeArea(pizzaLarge), ",cost per cm2 ", pizzaLarge.CostPerCm2())

	fmt.Println("Rectangle area", CulcShapeArea(rectangle))
}
