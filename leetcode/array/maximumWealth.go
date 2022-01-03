package main

import (
	"fmt"
)

// Input: accounts = [[1,2,3],[3,2,1]]
// Output: 6
// Explanation:
// 1st customer has wealth = 1 + 2 + 3 = 6
// 2nd customer has wealth = 3 + 2 + 1 = 6
// Both customers are considered the richest with a wealth of 6 each, so return 6

func main() {
	array := [][]int{{1,5},{7,3},{3,5}}

	value := maximumWealth(array)

	fmt.Println(value)

}

func maximumWealth(accounts [][]int) int {

	result := 0

	for _, v := range accounts {
		allValue := 0
		for _, v_in := range v {
			allValue = allValue + v_in
		} 	

		if (allValue >= result) {
			result = allValue
		}
	}

	return result
}