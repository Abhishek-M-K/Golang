package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza app !"
	println(welcome)

	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza, ", input)

	numRates, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if err != nil {
		fmt.Println(err)
		//panic(err) ends the program
	} else {
		fmt.Println("Modified rating is ", numRates+1)
	}


}
