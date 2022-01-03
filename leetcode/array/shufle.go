package main

import(
	"fmt"
)

// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
func main() {

	array := []int{2,5,1,3,4,7}
	n := 3
	value := shuffle(array, n)
	fmt.Println(value)

}

func shuffle(nums []int, n int) []int {
    start_y := n
	array := []int{}
	for i := 0; i < n; i++ {
		array = append(array, nums[i])
		array = append(array, nums[start_y])
		start_y = start_y + 1
	}
	return array
}