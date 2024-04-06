package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (r circle) perim() float64 {
	return math.Pi * 2 * r.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 1, height: 12.432}
	c := circle{radius: 1.000001}
	measure(r)
	measure(c)
}
