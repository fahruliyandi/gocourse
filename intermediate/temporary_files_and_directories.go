package basic

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//crete temp file
	// tempFile, err := os.CreateTemp(".", "temporaryFile")
	// checkError(err)
	// fmt.Println(tempFile.Name())
	// defer os.Remove(tempFile.Name())
	// defer tempFile.Close()

	//create dir temp
	dirTemp, err := os.MkdirTemp(".", "goCourseTempDir")
	checkError(err)
	fmt.Println("Dir Temp :", dirTemp)
	os.RemoveAll(dirTemp)

}
