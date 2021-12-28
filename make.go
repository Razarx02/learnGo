package main

import "fmt"


func main() {
	var  x []int

	x1 := []int{}
	x2 := make([]int, 5)
	fmt.Println(x)
	fmt.Println(x1)
	fmt.Println(x2)

	x3 := append(x2, 40)
	fmt.Println(x3)

	x4 := make([]int, 3)

	copy(x3, x4)

	fmt.Println(x4)

}