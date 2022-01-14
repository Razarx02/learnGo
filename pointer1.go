package main

import(
	"fmt"
)


func main() {

	d := 2

	square(&d)

	fmt.Println(d)

}



func  square(x *int) {
	*x = *x * *x 
}