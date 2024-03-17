package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Here are the slices !")

	/*var foodList = []string{"Chicken Tikka Biryani", "Mutter Paneer", "Subway Sandwich"}
	fmt.Printf("Type of list if %T\n", foodList)

	foodList = append(foodList, "Dal Tadka")
	fmt.Println("My Food List : ", foodList)
	fmt.Println(foodList[1:]) 

	topScores := make([]int, 3) //initially we have memory allocated for 3 elements
	topScores[0] = 100
	topScores[1] = 98
	topScores[2] = 93*/

	//as soon as we append something and it exceeds the allocated memory, 
	//it will reallocate the memory and copy the elements to the new memory location
	/*topScores = append(topScores, 89) 
	fmt.Println("Top Scores : ", topScores)

	sort.Ints(topScores)
	fmt.Println("Sorted Top Scores : ", topScores)*/

	var myTechs = []string{"reactjs", "nextjs", "nodejs", "expressjs", "golang", "mongodb"}
	fmt.Println("My Techs : ", myTechs)

	var ind int = 2

	myTechs = append(myTechs[:ind], myTechs[ind+1:]...)
	fmt.Println("My Updated Techs : ", myTechs)
}
