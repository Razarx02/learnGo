package main

import (
	"fmt"
	"reflect"
)

func main() {
	value := 'A'
	value_1 := 40


	fmt.Println("int32 ", reflect.TypeOf(value))
	fmt.Println("int ", reflect.TypeOf(value_1) )
	// fmt.Println(value + value_1)  error
}