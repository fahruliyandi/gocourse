package basic

import "fmt"

func main() {

	sum := add(1, 2)
	fmt.Println(sum)

	greet := func() {
		fmt.Println("hello anonimous function")
	}

	greet()

	operation := add
	result := operation(5, 5)
	fmt.Println(result)

	res := ApplyOperation(10, 10, add)
	fmt.Println(res)

	multiplier := CreateMultiplier(2)
	fmt.Println(multiplier(6))

}

func ApplyOperation(a int, b int, fnc func(int, int) int) int {
	return fnc(a, b)
}

func add(a int, b int) int {
	return a + b
}

func CreateMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
