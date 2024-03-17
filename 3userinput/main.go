package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to GoLang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza : ")

	//comma ok  || err ok
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input, "Error : ", err)
	fmt.Printf("Type of input is %T", input)
}