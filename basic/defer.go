package basic

import "fmt"

func main() {
	process(100)
}

func process(i int) {
	defer fmt.Println("Nilai defer i", i)
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	defer fmt.Println("third defer")
	fmt.Println("normal")
	i++
	fmt.Println(i)

}
