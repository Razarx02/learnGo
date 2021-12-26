package main

import "fmt"

var rec = 1024
func main() {
	rec := "las"
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v from %v \n", i, 10)
	}
	fmt.Println("First loop is finished")
	fmt.Println(rec)
	PrintRec()
	PrintRec_1(rec)
}

func PrintRec() {
	fmt.Println(rec)	
}

func PrintRec_1(rec string ) {
	fmt.Println(rec)	
}

// fmt.Println(rec)