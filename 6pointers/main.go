package main

import "fmt"

func main() {
	fmt.Println("Here are the pointers !")

	//var ptr *int //pointer holding the address of an int , default value is nil
	// fmt.Println("Value of ptr : ", ptr)

	myNum := 23

	var ptr  = &myNum

	fmt.Println("Value of pointer is : ", ptr)   //it is a reference to the memory location of myNum
	fmt.Println("Value of pointer is : ", *ptr)	//dereferencing the pointer to get the value at the memory location

	*ptr=*ptr+2
	fmt.Println("Value of update myNum is : ", myNum )	//myNum is updated as the pointer is updated
}
