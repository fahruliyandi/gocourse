package basic

import "fmt"

func main() {

	//arrray di go fixed size
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	fruits := [4]string{"apple", "banana", "orange", "grapes"}
	fmt.Println(fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[2] = 200
	fmt.Println(originalArray)
	fmt.Println(copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("index : %d, value : %d\n", index, value)
	}

	//underscrore untuk store unused value
	for _, value := range numbers {
		fmt.Printf("value : %d\n", value)
	}

	//comparing arrays
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array1 == array2)

	//multidimension array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	//pass by reference
	var originalArray1 [3]int = [3]int{
		1, 2, 3,
	}

	var copiedArray1 *[3]int
	copiedArray1 = &originalArray1
	copiedArray1[2] = 200

	fmt.Println(originalArray1)
	fmt.Println(copiedArray1)

}
