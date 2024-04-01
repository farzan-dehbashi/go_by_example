package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

func (r *rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getVatar() float64 {
	return math.Sqrt(r.width*r.width + r.height*r.height)
}

func main() {
	r := rectangle{width: 1, height: 1}
	fmt.Println("area -> ", r.getArea())
	fmt.Println("vatar ->", r.getVatar())

	rp := &r
	fmt.Println(rp.getArea())
	fmt.Println(rp.getVatar())

}
