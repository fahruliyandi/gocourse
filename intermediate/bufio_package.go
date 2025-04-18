package basic

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("Heloo, bufio package!\nhow are you doing"))
	// //reading the data into byte slice
	// data := make([]byte, 20)
	// ret_value, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("error reading data", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes : %s\n", ret_value, data[:ret_value])

	// ret_value_str, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("error reading string", err)
	// 	return
	// }
	// fmt.Println("reading string :", ret_value_str)

	//buffio writer
	writer := bufio.NewWriter(os.Stdout)

	//writing byte slice
	data := []byte("Hello, buffer io package")
	result, err := writer.Write(data)
	if err != nil {
		fmt.Println("error reading string", err)
		return
	}
	fmt.Printf("We wrote %d bytes\n", result)

	// flush the buffer to ensure all data is wrtten to stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("error flashing", err)
		return
	}

	//writing string
	str := "this is string\n"
	n, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("error write string", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("error writing", err)
		return
	}

}
