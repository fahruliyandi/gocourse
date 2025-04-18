package basic

import (
	"fmt"
	"path/filepath"
)

func main() {
	relativePath := `gocourse\example.txt`
	absolutePath := `D:\Tutorial\gocourse\example.txt`

	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split("user/home/docs/file.txt")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
	filePathBase := filepath.Base("user/home/docs/file.txt")
	fmt.Println("file path base:", filePathBase)

	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))

	//convert relative path ke absolute path

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error converting relative to abs path")
		return
	}

	fmt.Println(absPath)

}
