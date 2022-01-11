package main

import(
	"fmt"
)


// Input: s = "(1+(2*3)+((8)/4))+1"
// Output: 3
// Explanation: Digit 8 is inside of 3 nested parentheses in the string.

// Input: s = "(1)+((2))+(((3)))"
// Output: 3

func main() {

	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))
}

func maxDepth(s string) int {
	// left скобки с левой стороны, right - с правой
	left := 0
	right := 0
	result := 0
	
	for indexOut, vOut := range s {
		sumeDepth := 0 // Количество скобки
		if(vOut >= '0' && vOut <= '9' || vOut == '+' || vOut == '-' || vOut == '(' || vOut == ')') {
			for indexIn, vIn := range s { 
				if(indexIn < indexOut) {
					if(vIn == '(') {
						left++
					} 
					if(vIn == ')') {
						left--
					} 
				} else {
					if(vIn == ')') {
						right++
					}
					if(vIn == '(') {
						right--
					}
				}
			} 
			
			if(left <= right && left != 0 && right != 0 ) {
				sumeDepth = left
			} else if(left >= right && left != 0 && right != 0) {
				sumeDepth = right
			}
		}

		if(result <= sumeDepth) {
			result = sumeDepth
		}
		
		left = 0
		right = 0
	}

	return result

}