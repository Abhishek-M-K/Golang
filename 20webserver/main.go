package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web request handling in GoLang")
	fmt.Println("--------------------------------GET Request--------------------------------")
	GetRequest("http://localhost:3000/")
	fmt.Println("--------------------------------POST Request--------------------------------")
	PostRequest("http://localhost:3000/manga")
}

func GetRequest(url string) {
// 	const url = "http://localhost:3000/"

	res, err := http.Get(url)

	CheckNilError(err)

	defer res.Body.Close()

	fmt.Println("Status:", res.Status)
	// fmt.Println("Header:", res.Header)
	
	var resString strings.Builder  // This is a string builder to store the response content
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)

	fmt.Println("Content Length:", byteCount)
	fmt.Println("Content:", resString.String())

	// fmt.Println("Content:", string(content))  // This is another way to print the content
	// fmt.Println("Content Length:", res.ContentLength)

}

func PostRequest(url string) {
	//dummy json payload
	reqBody := strings.NewReader(`
		{
			"mangaName" : "Berserk",
			"genre" : "Action",
			"author" : "Kentaro Miura"
		}
	`)

	res, err := http.Post(url, "application/json", reqBody)

	CheckNilError(err)

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Status:", res.Status)
	fmt.Println(string(content))
}

func CheckNilError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
