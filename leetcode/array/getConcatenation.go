package main

import(
	"fmt"
)

func main() {

	value := getConcatenation([]int{1,2,3})
	fmt.Println(value)

}

func getConcatenation(nums []int) []int {
	array_1 := nums
	for _, i := range nums {
		array_1 = append(array_1, i)
	}	
	return array_1
}