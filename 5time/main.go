package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study with GoLang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))

	customTime := time.Date(2003, time.May, 29, 13, 0, 0, 0, time.UTC)
	fmt.Println("My custom date : ",customTime.Format("01-02-2006 Monday"))
}
