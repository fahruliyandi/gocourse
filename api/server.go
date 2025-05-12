package api

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello from server")
	})

	const serverAddr string = "127.0.0.1:3000"

	fmt.Println("Server is listening port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}
