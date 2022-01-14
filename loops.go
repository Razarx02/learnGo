package main

import "fmt"

func main() {
	// For can worl Like while
	counter := 0

	for {
		if counter > 10 {
			return
		} else {
			counter++
		}
		fmt.Print(counter)
	}

}