package main

import "fmt"

func welcomeMsg(){
	fmt.Println("Namaste India !")
}

func multi(a int, b int) int {
	return a * b
}

func proMulti(values ...int) (int,string) {
	total := 1

	for _, value := range values {
		total *= value
	}

	return total, "This is the result of proMulti function"
}

func main() {
	welcomeMsg()
	fmt.Println("Let's learn about functions in Go!")

	// defining functions inside functions are not allowed

	result := multi(10,4)
	fmt.Println("Result of multi function is --> ", result)

	answer, _ := proMulti(2,5,4);
	fmt.Println("Result of proMulti function is --> ", answer)

}
