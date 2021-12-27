package main

import "fmt"

func main() {
	
	count := "Hello world"
	text := ""

	for _,i := range count {
		fmt.Println(i)
		text = text + string(i)
	}

	fmt.Println(text)
}


