//interface segregation
//jangan kebanyakan method di interface nanti struct terpaksa implementnya

package SOLID

import "fmt"

//contoh jelek
// type machine interface{
// 	print()
// 	scan()
// 	fax()
// }

// //old printer cuma bisa print tapi dipaksa impelement scan sama fax, solusinya buat interface pisah pisah
// type oldPrinter struct{}
// func(m *oldPrinter) print(){
// 	fmt.Println("Ngeprint")
// }
// func(m *oldPrinter) scan(){
// 	fmt.Println("ga bisa print")
// }
// func(m *oldPrinter) fax(){
// 	fmt.Println("ga bisa fax")
// }

type printer interface {
	print()
}

type scanner interface {
	scan()
}

type faxer interface {
	fax()
}

type oldPrinter struct{}

func (old *oldPrinter) print() {
	fmt.Println("Ngeprint dari old printer")
}

type modernPrinter struct{}

func (mp *modernPrinter) print() {
	fmt.Println("Print dari modern printer")
}

func (mp *modernPrinter) scan() {
	fmt.Println("Scan dari modern printer")
}

func (mp *modernPrinter) fax() {
	fmt.Println("Fax dari modern printer")
}

func main() {
	op := &oldPrinter{}
	mp := &modernPrinter{}

	op.print()

	mp.print()
	mp.scan()
	mp.fax()
}
