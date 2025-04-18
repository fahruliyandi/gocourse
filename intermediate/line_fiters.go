package basic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer func() {
		fmt.Println("Closing file...")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	keyword := "importan"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Println("Filtered Line:", updatedLine)
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning", err)
		return
	}
}
