// open closed
// nambah boleh tapi modif koding yang udah ada ga boleh
package SOLID

import "fmt"

//contoh ga pake open closed
func totalBill(member string, total int) int {
	diskon := 0
	if member == "SILVER" {
		diskon = 10
		return total - diskon
	} else if member == "PLATINUM" {
		diskon = 20
		return total - diskon
	} else {
		return total
	}
	//kalau kita mau nambah member GOLD harus modif function totalBill ini
}

//contoh pake open closed principle
type Diskon interface {
	hitung(total int) int
}

type Silver struct{}

func (s *Silver) hitung(total int) int {
	return total - 10
}

type Platinum struct{}

func (p *Platinum) hitung(total int) int {
	return total - 20
}

func prosesBayar(d Diskon, total int) int {
	return d.hitung(total)
}

//kalau mau tambah member GOLD tinggal tambah struct GOLD

func main() {

	// fmt.Println("jumlah yang harus dibayarkan : ", totalBill("SILVER", 100))
	p := &Silver{}
	fmt.Println("Total yang harus dibayarkan :", prosesBayar(p, 1000))
}
