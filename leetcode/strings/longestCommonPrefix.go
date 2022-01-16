package main 

import (
	"fmt"
)


func main() {

	strs := []string{"flower","flow","flight"}	

	result := longestCommonPrefix(strs)
	
	fmt.Println(result)


}


func longestCommonPrefix(strs []string) string {
    if(len(strs) == 1) {
		return strs[0]
	}
	minIndex := minElementIndex(strs) 
	text := ""
	check := false
	for i, v := range strs[minIndex] {
		for indexArray, valueArray := range strs {
			if (indexArray != minIndex) {
				if(rune(valueArray[i]) == v) {
					check = true
				} else {
					check = false
					return text
				}
			}
		}

		if(check) {
			text = text + string(v)
		}
	}

	return text

}

func minElementIndex(strs []string) int {
	index := 0
	size := len(strs[0])
	for i, str := range strs {
		if(len(str) <= size) {
			size = len(str)
			index = i
		}
	}

	return index
}