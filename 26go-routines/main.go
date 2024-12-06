package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup // WaitGroup is a struct that allows us to wait for a collection of goroutines to finish
// pointer

var mutex sync.Mutex // Mutex is a struct that allows us to lock and unlock a block of code
// pointer

var signals = []string{"test"}

func main(){
	// go greeter("Hello")
	greeter("World")

	webSites := []string{
		"https://google.com",
		"https://go.dev",
	}

	for _, url := range webSites {
		go getStatusCode(url)
		waitGroup.Add(1) // increment the waitGroup counter
	}

	waitGroup.Wait() // wait for all goroutines to finish

	fmt.Println(signals)
}

func greeter(s string){
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(url string){
	defer waitGroup.Done() // decrement the waitGroup counter

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in http request", err)
		return	
	}

	mutex.Lock()
	signals = append(signals, url)
	mutex.Unlock()
	fmt.Println("Status code for", url, "is", res.StatusCode)
	// fmt.Printf("%d status code for %s", res.StatusCode, url)
}