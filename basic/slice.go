package basic

import (
	"fmt"
	"slices"
)

func main() {

	//slice sama kaya array cuma dinamic size
	// var numbers []int
	// var numbers2 = []int{1, 2, 3}
	// numbers3 := []int{1, 2, 3}
	// slice = make([]int, 5)
	var a = [5]int{1, 2, 3, 4, 5}
	var slice1 = a[1:4]
	fmt.Println(slice1)
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	//copy slice
	var slice2 = make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice2)

	//null slice
	// var nilSlice []int

	for i, v := range slice2 {
		fmt.Println(i, v)
	}

	// slice2[3] = 100
	// fmt.Println(slice2)

	if slices.Equal(slice1, slice2) {
		fmt.Println("Slices are equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		var innerLen int = i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	var slice3 = slice1[1:3]
	fmt.Println(slice3)
	fmt.Println(cap(slice3))
	fmt.Println(len(slice3))

	var slice4 []int = []int{1, 3, 4, 5}
	var slice5 = slice4
	slice5[0] = 500
	fmt.Println(slice4)
	fmt.Println(slice5)

	//slice sering digunakan daripada array
	//slice auto pass by reference

}
