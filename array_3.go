package main

import (
	"fmt"
	"reflect"
)


func main() {
	arrayStr := [3]string{}
	fmt.Println(arrayStr)
	fmt.Println(reflect.TypeOf(arrayStr))
}