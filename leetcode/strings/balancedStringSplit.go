package main

import(
	"fmt"
)


// Balanced strings are those that have an equal quantity of 'L' and 'R' characters.
// Given a balanced string s, split it in the maximum amount of balanced strings.
// Return the maximum amount of split balanced strings.

// Input: s = "RLRRLLRLRL"
// Output: 4
// Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

// Input: s = "RLLLLRRRLR"
// Output: 3
// Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

// Input: s = "LLLLRRRR"
// Output: 1
// Explanation: s can be split into "LLLLRRRR".


func main() {
	s := "RLLLLRRRLR"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
    
	
	balance := 0
	splits := 0
	for _, v := range s {
		
		if(v == 'L') {
			balance--
		} else {
			balance++
		}

		if(balance == 0) {
			splits++
		}
	}

	return splits

}