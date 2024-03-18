package main

import "fmt"

func main() {
	fmt.Println("Lets learn loops !")

	techs := []string{"react","nextjs","node","flutter","aws"}
	// fmt.Println(techs)

	//for loop
	for i := 0; i < len(techs); i++ {
		fmt.Println(techs[i])
	}

	//for range loop
	for i := range techs {
		fmt.Println("Using range : ", techs[i])
	}

	//for each loop
	for i, tech := range techs {
		fmt.Printf("At index %v , the tech is %v\n ", i, tech)
	}

	//rogue value similar to while loop
	rougeValue := 5
	for rougeValue < 10 {
		// if rougeValue == 7 {
		// 	break
		// }
		if rougeValue == 10 {
			goto destiny
		}

		fmt.Println("Rouge value is ", rougeValue)
		rougeValue++
	}

	destiny:
		fmt.Println("I made it !")
}
