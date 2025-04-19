package intermediate

import "fmt"

func main() {
	//numeric type convertion
	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 3.14
	f := int(e)
	fmt.Println(c, f)

	g := "heloo"
	h := []byte(g)
	fmt.Println(h)
	i := []byte{104, 101, 108, 111, 111}
	j := string(i)
	fmt.Println(j)

}
