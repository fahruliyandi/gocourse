//Liskov Substitution Principle (LSP)
//Semua struct yang implement interface harus bisa digunakan
// di mana pun interface itu dipakai, tanpa bikin program
//  jadi error atau berperilaku aneh

package SOLID

import "fmt"

// type burung interface {
// 	terbang()
// }

// type elang struct{}
// func(e *elang) terbang(){
// 	fmt.Println("Elang terbang")
// }

// //contoh LSP yang buruk karena meskipun ayam itu burung cuma ayam ga bisa terbang, solusinya buat 2 interface
// type ayam struct{}
// func (a *ayam) terbang(){
// 	fmt.Println("Ayam ga bisa terbang")
// }

//solusi
type burung interface {
	berkicau()
}

type bisaTerbang interface {
	terbang()
}

type elang struct{}

func (e *elang) berkicau() {
	fmt.Println("Elang berkicau")
}
func (e *elang) terbang() {
	fmt.Println("Elang terbang")
}

type ayam struct{}

func (a *ayam) berkicau() {
	fmt.Println("Ayam berkicau")
}

func main() {

}
