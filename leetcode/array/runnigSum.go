package main

import (
	"fmt"
)

// Input: nums = [1,1,1,1,1]
// Output: [1,2,3,4,5]
// Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1]

func main() {

	array := []int{1,2,3,4}
	result:= runningSum(array)
	fmt.Println(result)
}


func runningSum(nums []int) []int {
    
	result := []int{}
	for i := 0; i < len(nums); i++ {
		value := 0
		for n := 0; n <= i; n++ {
			value = value + nums[n]
		}
		result = append(result, value)
	}

	return result
}