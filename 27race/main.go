package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Printf("Hello, World!\n")

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	var scores = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mt *sync.Mutex){
		fmt.Println("One R")
		mt.Lock()	
		scores = append(scores, 1)
		mt.Unlock()
		wg.Done()
	}(&wg, &mutex)

	go func(wg *sync.WaitGroup, mt *sync.Mutex){
		fmt.Println("Two R")
		mt.Lock()
		scores = append(scores, 2)
		mt.Unlock()
		wg.Done()
	}(&wg, &mutex)

	go func(wg *sync.WaitGroup, mt *sync.Mutex){
		fmt.Println("Three R")
		mt.Lock()
		scores = append(scores, 3)
		mt.Unlock()
		wg.Done()
	}(&wg, &mutex)

	wg.Wait()	

	fmt.Println("Scores: ", scores)

	// status 66 for race condition
}