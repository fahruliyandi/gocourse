package basic

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// err := os.Mkdir("subdir", 0755)
	// CheckError(err)

	// defer func() {
	// 	fmt.Println("removing directory...")
	// 	//remove dir
	// 	os.RemoveAll("subdir")
	// }()

	//create dir sama filenya
	// os.WriteFile("subdir/file.txt", []byte("Hello"), 0755)

	// create nested dir
	// err := os.MkdirAll("subdir/parent/child", 0755)
	// CheckError(err)

	result, err := os.ReadDir(`subdir/parent`)
	CheckError(err)
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())

	}

	// entry, err := os.ReadDir(".")
	// CheckError(err)
	// for _, ent := range entry {
	// 	fmt.Println(ent.Name())
	// }

	// fmt.Println(len(entry))

	//change directory
	// err = os.Chdir("subdir/parent")
	CheckError(err)

	entry, err := os.ReadDir(".")
	CheckError(err)
	for _, ent := range entry {
		fmt.Println(ent.Name())
	}

	//get working directory
	// entry, err = os.ReadDir("../../..")
	pwd, err := os.Getwd()
	fmt.Println(pwd)

	//filepath.walk and filepath.walkdir
	pathfile := "subdir"
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("error walkdir")
			panic(err)
		}
		fmt.Println(path)
		return nil
	})
	CheckError(err)

	//remove dir
	err = os.RemoveAll("subdir")
	CheckError(err)
}
