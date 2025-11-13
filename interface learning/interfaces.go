package main

import "fmt"

type shapes interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width + r.height
}

func measure(s shapes) {
	fmt.Println("Area", s.area(), "Perimeter", s.perimeter())
}

func main() {

	rect := rectangle{
		width:  2.0,
		height: 5.0,
	}

	measure(rect)

	fmt.Println("this is the interfaces learning area")
}
