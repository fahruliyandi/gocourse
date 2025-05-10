package advance

import (
	"fmt"
	"sort"
)

type Person struct {
	nama string
	umur int
}

type By func(p1 *Person, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1 *Person, p2 *Person) bool
}

func (p *personSorter) Len() int {
	return len(p.people)
}

func (p *personSorter) Less(i int, j int) bool {
	return p.by(&p.people[i], &p.people[j])
}

func (p *personSorter) Swap(i int, j int) {
	p.people[i], p.people[j] = p.people[j], p.people[i]
}

func (by By) sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// type ByAge []Person

// func (p ByAge) Len() int {
// 	return len(p)
// }

// func (p ByAge) Less(i int, j int) bool {
// 	return p[i].umur < p[j].umur
// }

// func (p ByAge) Swap(i int, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

func main() {
	people := []Person{
		{nama: "Fahrul", umur: 36},
		{nama: "Annisa", umur: 35},
		{nama: "Fahira", umur: 3},
	}

	Umur := func(p1 *Person, p2 *Person) bool {
		return (*p1).umur < (*p2).umur
	}

	Name := func(p1 *Person, p2 *Person) bool {
		return (*p1).nama < (*p2).nama
	}

	By(Umur).sort(people)
	fmt.Println(people)
	By(Name).sort(people)
	fmt.Println(people)
	// 	sort.Sort(ByAge(people))
	// 	fmt.Println(people)
}
