package basic

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	width  float64
	length float64
}

type Rectangle1 struct {
	width  float64
	length float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (r Rectangle1) area() float64 {
	return r.width * r.length
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) perim() float64 {
	return 2 * (r.length * r.width)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := Rectangle{
		width:  10,
		length: 11,
	}

	c := Circle{
		radius: 100,
	}

	// r1 := Rectangle1{
	// 	width:  10,
	// 	length: 11,
	// }

	measure(r)
	measure(c)
	// measure(r1)

	myPrinter(1, "fahrul", true)

}

// function yang bisa terima semua tipe data
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
