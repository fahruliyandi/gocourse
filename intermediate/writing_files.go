package basic

import (
	"fmt"
	"os"
)

func main() {
	//create file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file :", err)
		return
	}
	defer file.Close()

	//write data to file
	data := []byte("Hello world\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file : ", err)
		return
	}

	fmt.Println("Data has been written to file sucessfuly")

	file, err = os.Create("WriteString.txt")
	if err != nil {
		fmt.Println("Error creating to file : ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello go")
	if err != nil {
		fmt.Println("Error writing to file : ", err)
		return
	}

	fmt.Println("Data has been written to writing file")

}
