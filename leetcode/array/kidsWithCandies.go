package main

import(
	"fmt"
)

// Input: candies = [2,3,5,1,3], extraCandies = 3
// Output: [true,true,true,false,true] 
// Explanation: If you give all extraCandies to:
// - Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
// - Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
// - Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
// - Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
// - Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.

func main() {

	array := []int{2,3,5,1,3}
	
	extraCandies := 3
	
	value := kidsWithCandies(array, extraCandies)
	
	fmt.Println(value)
}


func kidsWithCandies(candies []int, extraCandies int) []bool {

	array := make([]bool, len(candies))
	
	max_element, _ := maxMin(candies)

	for i,  v := range candies {
		if(v + extraCandies >=  max_element) {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	return array
}


func maxMin(array []int) ( int, int) {
	max := array[0]
	min := array[0]
	for _, v := range array {
		if (v >= max) {
			max = v
		}

		if (v <= min) {
			min = v
		}
	}
	return max, min
}