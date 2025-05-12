package api

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error making get requst", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
	}

	fmt.Println(string(body))

}
