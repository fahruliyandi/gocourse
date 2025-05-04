package latihan

// func ubahNilai(p *int) {
// 	*p++
// }

// type mahasiswa struct {
// 	nama string
// 	umur int
// }

// func ubahLangsung(m mahasiswa) {
// 	m.umur = 100
// }

// func ubahPointer(m *mahasiswa) {
// 	m.umur = 99
// }

// func ubahSlice(sl *[]int) {
// 	*sl = append(*sl, 100)
// }

type Person struct {
	nama string
	umur int
}

func newPerson(nama string, umur int) Person {
	return Person{
		nama: nama,
		umur: umur,
	}
}

func latihan() {

	// num := []int{1, 2, 3}
	// ubahSlice(&num)
	// fmt.Println(&num)

	// a := 10
	// p := &a

	// fmt.Println("Nilain a", a)
	// fmt.Println("Alamat nilai a", p)
	// fmt.Println("Nilai dari pointer p", *p)
	// *p = 20
	// fmt.Println("nilai a setelah dirubah pointer", a)

	// a := 10
	// ubahNilai(&a)
	// fmt.Println(a)

	// m1 := mahasiswa{
	// 	nama: "fahrul",
	// 	umur: 36,
	// }

	// ubahLangsung(m1)
	// fmt.Println(m1.umur)

	// m2 := &mahasiswa{
	// 	nama: "annisa",
	// 	umur: 35,
	// }

	// ubahPointer(m2)
	// fmt.Println(m2.umur)
}
