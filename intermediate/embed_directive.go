package basic

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basic
var basicFolder embed.FS

func main() {
	fmt.Println("embedded content", content)
	content, err := basicFolder.ReadFile("basic/hello.txt")
	if err != nil {
		fmt.Println("error read file", err)
	}
	fmt.Println(string(content))

	err = fs.WalkDir(basicFolder, "basic", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
