package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 6}
	rectangle := Rectangle{width: 5, height: 3}
	fmt.Println("Circle area ", getArea(circle))
	fmt.Println("Rectangle area ", getArea(rectangle))

}
