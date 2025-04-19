package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader :", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing from writer :", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error Closing :", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer")
	fmt.Println(buf.String())

}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("world")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader")
	}
	fmt.Println(buf.String())

}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("error opening or creating file:", err)
	}
	defer closeResource(file)

	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln("error opening or creating file:", err)
	}

}

func main() {

	fmt.Println("==Read from reader==")
	readFromReader(strings.NewReader("Hello reader"))

	fmt.Println("==Write to writer==")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writer")
	fmt.Println(writer.String())

	fmt.Println("==Buffer example==")
	bufferExample()

	fmt.Println("==Multi Reader==")
	multiReaderExample()

	fmt.Println("==Piple Example==")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "HEloo")

}
