package main

import (
	"fmt"
)


func main(){

	str := "1.1.1.1"

	result := defangIPaddr(str)

	fmt.Println(result)
}

func defangIPaddr(address string) string {
	 result := ""

	for _, v := range address {
		if( v == '.' ) {
			result = result + "[.]"
		} else {
			result = result + string(v)
		}
	}

	return result
}