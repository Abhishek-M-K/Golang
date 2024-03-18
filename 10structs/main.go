package main

import "fmt"

func main() {
	fmt.Println("Lets learn Structs !")

	// no super or parent class that means no inheritance

	abhishek := User{"Abhishek", "khandareabhishek98@gmail.com","9967198279", 20, true}
	fmt.Println(abhishek)
	fmt.Printf("Details of user Abhishek are : %+v\n", abhishek)
	fmt.Println("Phone number of Abhishek is : ", abhishek.Phone)
}

type User struct {
	Name string
	Email string
	Phone string
	Age int
	Active bool
}
