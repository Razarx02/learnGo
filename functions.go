package main

import "fmt"

func main() {
	value := GiveZero()
	fmt.Println(value)
}


func GiveZero() int {
	fmt.Println("Zero function")
	return 0
}