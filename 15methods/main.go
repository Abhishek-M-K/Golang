package main

import "fmt"

type User struct {
	Name string
	Email string
	Phone string
	Age int
	Active bool
}

func (u User) GetStatus(){
	fmt.Println("Is user active :", u.Active)
}

func (u User) UpdatePhone(){
	u.Phone = "1234567890"
	fmt.Println("Updated phone number is : ", u.Phone)
}

func main() {
	fmt.Println("Let's learn methods in Go!")

	//similar to methods

	abhishek := User{"Abhishek", "khandareabhishek98@gmail.com","9967198279", 20, true}
	fmt.Println(abhishek)
	abhishek.GetStatus()
	abhishek.UpdatePhone()
	fmt.Println(abhishek)

}
