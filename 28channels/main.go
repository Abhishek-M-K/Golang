package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning Go Channels")

	myChannel := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Read from channel
	go func (chh <- chan int, wg *sync.WaitGroup) {
		val, isChhOpen := <- myChannel

		if isChhOpen {
			fmt.Println(isChhOpen)
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed")
		}

		// fmt.Println(<- myChannel)
		fmt.Println("First go routine")
		wg.Done()
	}(myChannel, &wg)

	// Write to channel
	go func (chh chan <- int, wg *sync.WaitGroup) {
		close(myChannel) // close the channel
		// myChannel <- 12
		// myChannel <- 13
		fmt.Println("Second go routine")
		wg.Done()
	}(myChannel, &wg)

	// myChannel <- 42

	fmt.Println(<-myChannel) // deadlock // channels cannot be accessed without a goroutine
}

// Channels can be bidierectional or unidirectional
// <- chan int // read only channel
// chan <- int // write only channel