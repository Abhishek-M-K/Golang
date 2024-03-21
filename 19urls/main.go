package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://abhishekkhandare.me"

func main() {
	fmt.Println("Let's learn about Web Requests in Go!")
	fmt.Println("URL: ", myUrl)

	//parse
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Query: ", result.RawQuery)
	fmt.Println("Fragment: ", result.Fragment)
	fmt.Println("User: ", result.User)

	params := result.Query()
	fmt.Printf("Type: %T\n ", params)
	fmt.Println(params) //map[string][]string map[utm_source:[google] utm_medium:[cpc]] key value pair

	//construct a URL
	newUrl := &url.URL{
		Scheme:   "https",
		Host:    "abhishekkhandare.me",
		Path:    "/contact",
	}

	fmt.Println("New URL : ",newUrl.String())
}
