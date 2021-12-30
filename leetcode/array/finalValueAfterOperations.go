package main

import (
	"fmt"
)

func main() {
	result := finalValueAfterOperations([]string{"X++","--X", "--X"})
	fmt.Println(result)

}

func finalValueAfterOperations(operations []string) int {

	result := 0
	for _, v := range operations {
		value, err := containsCheck(v)
		if (err == true) {
			return 0
		} else if (value == true) {
			result++		
		} else if (value == false) {
			result--
		}
	}
	return result
	
}

func containsCheck(str string) (bool, bool) {
	for _, v := range str {
		if(v == '+') {
			return true, false
		} else if (v == '-') {
			return false, false
		} 
	}

	return false, true
}