package main

import (
	"fmt"
	"math"
)

type (
	Shape interface {
		Area() float64
	}

	Rectangle struct {
		Width  float64
		Height float64
	}
	Circle struct {
		Radius float64
	}
	Triangle struct {
		Base   float64
		Height float64
	}
)

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 4, Height: 3},
	}
	total := CalculateTotalArea(shapes)
	fmt.Println("Total area of all shapes:", total)

	c := Circle{
		Radius: 5,
	}

	r := Rectangle{
		Width: 10, Height: 5,
	}

	t := Triangle{
		Base: 4, Height: 3,
	}

	fmt.Println(c.Area() + r.Area() + t.Area())
	fmt.Println(c.Area())
	fmt.Println(r.Area())
	fmt.Println(t.Area())

}
