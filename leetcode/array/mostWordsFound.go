package main 

import(
	"fmt"
	"strings"
)

// Input: sentences = ["please wait", "continue to fight", "continue to win"]
// Output: 3
// Explanation: It is possible that multiple sentences contain the same number of words. 
// In this example, the second and third sentences (underlined) have the same number of words.

func main() {
	array := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	value := mostWordsFound(array)
	fmt.Println(value)
}

func mostWordsFound(sentences []string) int {
    
	value := 0
	for _, v := range sentences {
		size := len(strings.Split(v, " "))
		if( value <= size) {
			value = size
		}
	}
	return value
}