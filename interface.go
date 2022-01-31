package main

import (
	"fmt"
)

const (
	PI = 3.1416
)

type Shape interface {
	area() float64
	description() string
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) description() string {
	return "Rectangle"
}

func (c *Circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c *Circle) description() string {
	return "Circle"
}

func main() {

	r1 := Rectangle{10, 5}
	c1 := Circle{5.5}

	r2 := []Rectangle{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}

	c2 := []Circle{{4.5}, {6.5}, {3}, {7.9}}

	s1 := []Shape{}

	for i := 0; i < len(r2); i++ {
		s1 = append(s1, &r2[i])
	}

	for i := 0; i < len(c2); i++ {
		s1 = append(s1, &c2[i])
	}

	fmt.Println(s1)

	for _, s := range s1 {
		a := fmt.Sprintf("Area of %v : %v \n", s.description(), s.area())
		fmt.Println(a)

	}

	fmt.Println("###################")

	shapes := []Shape{&c1, &r1, &c1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(shape.description())
	}

}
