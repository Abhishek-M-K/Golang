package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://abhishekkhandare.me"

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Let's learn about Web Requests in Go!")

	res, err := http.Get(url)

	checkNilError(err)

	// fmt.Println("Response: ", res)

	fmt.Printf("Type : %T\n", res) // Type : *http.Response (Pointer to http.Response)
	fmt.Println("Status: ", res.Status)

	defer res.Body.Close()

	webContent, err := ioutil.ReadAll(res.Body)

	checkNilError(err)

	fmt.Println("Content: ", string(webContent))
}
