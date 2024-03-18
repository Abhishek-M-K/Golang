package main

import "fmt"

func main() {
	fmt.Println("Lets learn Maps !")

	languages := make(map[string]string)

	languages["TS"] = "Typescript"
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println("GO is shorthand for :", languages["GO"])

	delete(languages, "PY")
	fmt.Println("Updated list of languages : ",languages)

	// Iterating over the map
	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// comma ok helps to ignore any value
}
