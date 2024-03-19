package main

import "fmt"

func myDefer() {
	for i :=0 ; i<5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println("Let's learn defer in Go!")

	// LIFO logic is behind defer

	defer fmt.Println("Abhishek") // executes at the last of the function in which it is defined

	fmt.Println("Khandare")

	myDefer()
}
