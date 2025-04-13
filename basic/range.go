package basic

import (
	"fmt"
)

func main() {

	message := "hello world"
	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("index : %d, Rune : %c\n", i, v)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for i, _ := range numbers {
		if i == 0 {
			numbers[i] = 100
		}
	}

	fmt.Println(numbers)

	myMap := make(map[string]int)
	myMap["b"] = 2
	fmt.Println(myMap)

}
