package main

import(
	"fmt"
)

func main() {

	x := 10
	fmt.Println(x)
	change5(&x)
	fmt.Println(x)
	d := &x
	change15(d)
	fmt.Println(x)
}


func change5( variable *int) {
	*variable = 5
}


func change15( variable *int) {
	*variable = 15
}