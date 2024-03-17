package main

import "fmt"

func main() {
	fmt.Println("Here are the arrays !")

	var foodList [5]string
	foodList[0] = "Chinese Fried Rice"
	foodList[1] = "Chicken Tikka Biryani"
	foodList[2] = "Mutter Paneer"
	foodList[3] = "Dal Tadka"
	foodList[4] = "Subway Sandwich"

	fmt.Println("My Food List : ", foodList, "Number of items - ", len(foodList))

	var biscuits = [3]string{"Parle-G", "Bourbon", "Good-Day"}
	fmt.Println("Biscuits List : ", biscuits, "Number of items - ", len(biscuits)) 
}
