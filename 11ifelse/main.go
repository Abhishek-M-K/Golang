package main

import "fmt"

func main() {
	fmt.Println("Lets learn if else !")

	age := 35

	var answer string

	if age < 10 {
		answer = "Kid"
	} else if age < 20 {
		answer = "Teen"
	} else if age < 30 {
		answer = "Adult"
	} else if age > 30 && age < 40 {
		answer = "Middle Aged"
	} else {
		answer = "Old"
	}

	fmt.Println("You are a : ", answer)


	if 12%2 ==0 {
		fmt.Println("12 is even")
	}else{
		fmt.Println("12 is odd")
	}

	if num := 8; num<10 {
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is greater than 10")
	}

}
