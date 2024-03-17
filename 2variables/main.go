package main

import "fmt"

const LoginToken string = "asjdhjashdjkasdkjashd" //public variable


func main() {
	fmt.Println("Variables in GoLang")

	var userName string = "John Doe"
	fmt.Println("Value of variable : ", userName)
	fmt.Printf("Type of variable : %T \n", userName)

	var isAuthorized bool = false
	fmt.Println("Is user authorized ? : ", isAuthorized)

	var smallInt uint8 = 255 // 0 to 255 bound value
	fmt.Println(smallInt)
	fmt.Printf("Type of variable : %T \n", smallInt)

	var floatVal float32 = 3.14123123444473627
	fmt.Println(floatVal)
	fmt.Printf("Type of variable : %T \n", floatVal)

	//aliases
	var newOne int
	fmt.Println(newOne)
	fmt.Printf("Type of variable : %T \n", newOne)

	//implicit types
	var domain = "abhishekkhandare.me"
	fmt.Println(domain)
	fmt.Printf("Type of variable : %T \n", domain) //auto detects type of var based on value assigned

	//walrus operator (not allowed to use outside of the method) , no var keyword 
	age := 20
	fmt.Println(age)

	fmt.Println(LoginToken)

}