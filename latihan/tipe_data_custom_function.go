package latihan

import "fmt"

type mathOperation func(a int, b int) int

func (mo mathOperation) describe(action string) {
	fmt.Println("Ini adalah hasil operasi ", action)
}

func (mo mathOperation) callWithLog(a int, b int) int {
	fmt.Printf("Memanggil fungsi untuk hitung %d dan %d\n", a, b)
	result := mo(a, b)
	fmt.Println("result", result)
	return result
}

func latihan() {
	add := func(a int, b int) int {
		return a + b
	}

	mathOperation(add).describe("Penjumlahan")
	fmt.Println("hasil dari ", add(10, 10))

	fmt.Println(mathOperation(add).callWithLog(10, 10))
}

//====================================
// type mathOperation func(a int, b int) int

// func operation(action string) mathOperation {
// 	if action == "Add" {
// 		return func(a int, b int) int {
// 			return a + b
// 		}
// 	} else if action == "Substract" {
// 		return func(a int, b int) int {
// 			return a - b
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	add := operation("Add")
// 	subs := operation("Substract")

// 	fmt.Println(add(10, 10))
// 	fmt.Println(subs(10, 10))

// }

//===================================
// type callback func(str string)

// func performAction(action string, callback callback) {
// 	fmt.Println("Performing Action")
// 	callback(action)
// }

// func main() {
// 	logCallBack := func(str string) {
// 		fmt.Println("Action Complete", str)
// 	}

// 	performAction("Download File", logCallBack)
// }

//========================================
// type filter func(a int) bool

// func greaterThanTen(a int) bool {
// 	return a > 10
// }

// func lowerThanTen(a int) bool {
// 	return a < 10
// }

// func filteredNumbers(nums []int, filter filter) []int {
// 	var result []int
// 	for _, i := range nums {
// 		if filter(i) {
// 			result = append(result, i)
// 		}
// 	}

// 	return result
// }

// func main() {

// 	numbers := []int{10, 11, 7, 12, 15, 9}
// 	result := filteredNumbers(numbers, greaterThanTen)
// 	fmt.Println(result)
// 	result = filteredNumbers(numbers, lowerThanTen)
// 	fmt.Println(result)

// }

//=====================================
// type operation func(a int, b int) int

// func main() {

// 	var op operation

// 	add := func(a int, b int) int {
// 		return a + b
// 	}

// 	op = add
// 	fmt.Println(op(10, 10))

// 	substract := func(a int, b int) int {
// 		return a - b
// 	}

// 	op = substract

// 	fmt.Println(op(10, 10))

// }
