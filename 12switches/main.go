package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Lets learn switch !")

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	diceNum := rand.Intn(6)+1
	fmt.Println("You rolled a ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
	case 3:
		fmt.Println("You rolled a 3")
	case 4:
		fmt.Println("You rolled a 4")
	case 5:
		fmt.Println("You rolled a 5")
	case 6:
		fmt.Println("Wohoo, you rolled a 6 ! Roll again")
	default:
		fmt.Println("Invalid dice roll")
	}

}
