package basic

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println(result)

	result = a - b
	fmt.Println(result)

	result = a * b
	fmt.Println(result)

	result = a / b
	fmt.Println(result)

	result = a % b
	fmt.Println(result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	//overflow with signed integer
	//overflow ketika nilai lebih besar daripada yang bisa ditampung data typenya
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//overflow with signed integer
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	//underflow with floating point
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
