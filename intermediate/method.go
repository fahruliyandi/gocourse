package basic

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

//method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rec := Rectangle{
		length: 10,
		width:  9,
	}
	area := rec.Area()
	fmt.Println(area)
	rec.Scale(2)
	area = rec.Area()
	fmt.Println(area)

	num := myInt(1)
	num2 := myInt(-2)
	fmt.Println(num.isPositive())
	fmt.Println(num2.isPositive())
	fmt.Println(num.welcomeMassage())

	sp := Shape{
		Rectangle: Rectangle{
			length: 10, width: 11,
		},
	}

	fmt.Println(sp.Area())
}

//method on a user-defined type
type myInt int

func (m myInt) isPositive() bool {
	return m > 0
}

func (myInt) welcomeMassage() string {
	return " Welcome to myInt"
}
