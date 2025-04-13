package basic

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	substraction := func() func(int) int {
		countdown := 100
		return func(i int) int {
			countdown = countdown - i
			return countdown
		}
	}()

	//bukan closure
	substraction2 := func(i int) int {
		countdown := 100
		return countdown - i
	}

	fmt.Println(substraction(10))
	fmt.Println(substraction(10))
	fmt.Println(substraction(10))
	fmt.Println(substraction(10))
	fmt.Println(substraction(10))

	fmt.Println(substraction2(10))
	fmt.Println(substraction2(10))
	fmt.Println(substraction2(10))
}

func adder() func() int {
	i := 0
	fmt.Println("previous values of i :", i)
	return func() int {
		i++
		fmt.Println("Add one to i")
		return i
	}
}
