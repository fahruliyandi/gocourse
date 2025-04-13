package basic

import (
	"fmt"
)

func main() {
	//maps is unordered
	var myMap map[string]int = map[string]int{
		"key1": 1,
	}

	myMap["key2"] = 2
	fmt.Println(myMap)
	myMap["key1"] = 10
	fmt.Println(myMap)
	delete(myMap, "key1")
	fmt.Println(myMap)

	clear(myMap)
	fmt.Println(myMap)

	_, unknownvalue := myMap["key1"]
	fmt.Println(unknownvalue)

	var myMap2 map[string]int = map[string]int{
		"a": 1,
		"b": 2,
	}

	for key, value := range myMap2 {
		fmt.Println(key, value)
	}

	myMap3 := make(map[string]string)
	myMap3["a"] = "10"
	fmt.Println(len(myMap3))
	fmt.Println(myMap3)

	//nested map

}
