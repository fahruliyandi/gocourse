package basic

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating files :", err)
		return
	}
	// defer file.Close()

	_, err = file.WriteString("Hello world\n\nHOHOHO\n\n\n")
	if err != nil {
		fmt.Println("Error writing files :", err)
		return
	}

	fmt.Println("File created and written")

	file, err = os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening files :", err)
		return
	}
	defer func() {
		fmt.Println("Closing file...")
		file.Close()
	}()

	fmt.Println("File opened successfully")

	// //read the content the file
	// data := make([]byte, 1024)
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading content file:", err)
	// 	return
	// }
	// fmt.Println("File contet : ", string(data))

	scanner := bufio.NewScanner(file)
	//read content line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println("Line :", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
