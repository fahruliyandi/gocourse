package basic

import (
	"fmt"
	"net/url"
)

func main() {
	//[scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL", err)
		return
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=Jhon&age=30"
	parsedUrl1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing URL", err)
		return
	}

	queryParam := parsedUrl1.Query()
	fmt.Println(queryParam)
	fmt.Println("name : ", queryParam.Get("name"))
	fmt.Println("age : ", queryParam.Get("age"))

	//building url
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	query := baseUrl.Query()
	query.Set("name", "Jhon")
	query.Set("age", "30")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Base Url :", baseUrl.String())

	values := url.Values{}
	//add key value pair to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("country", "indo")

	encodedQuery := values.Encode()
	println(encodedQuery)

	//base url
	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Println(fullUrl)

}
