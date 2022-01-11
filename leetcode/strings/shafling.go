package main

import(
	"fmt"
	"sort"
 )


// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"
// Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.


func main() {
	s := "codeleet" 
	indices := []int{4,5,6,7,0,2,1,3}
	
	result := shuffling(s, indices)
	fmt.Println(result)

}


func shuffling(s string, indices []int) string {

	mymap := map[int]rune{}
	text := ""
	for i, v := range s {
		mymap[indices[i]] = v 
	}

	sort.Ints(indices)
	
	for _, v := range indices {
		text = text + string(mymap[v])
	}

	return text
}