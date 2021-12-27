package main

import (
	"fmt"
	"reflect"
)

func main() {

	array_1 := [10][2][]int{{{15}}, {{20},{30}}}
	fmt.Println(array_1[0][0][0])
	fmt.Println(array_1)
	fmt.Println("Array can be different")

	fmt.Println(reflect.TypeOf(array_1))

}