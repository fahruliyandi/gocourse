package basic

import "fmt"

func main() {
	//pointers adalah variable yang menyimpan memori address variable lainnya
	var ptr *int
	var a int = 10
	ptr = &a //referencing pointer
	fmt.Println(ptr)
	fmt.Println(&a)
	fmt.Println(*ptr) // untuk mendapatkan value dari pointer tersebut

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
