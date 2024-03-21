package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile(fileName string){
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("File content: ", string(dataByte))
}

func main() {
	fmt.Println("Let's learn File Handling in Go!")

	content := "Oda Eichiro is the best manga artist!"

	file, err := os.Create("./myFile.txt")

	checkNilError(err)

	// if err != nil {
	// 	panic(err) // similar to throw new Error() in JavaScript
	// }

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Printf("File created with %v characters\n", length)
	readFile(file.Name()) //defer file.Close()
	file.Close()
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
