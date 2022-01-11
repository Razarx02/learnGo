package main

import(
	"fmt"
)

// Input: s = "is2 sentence4 This1 a3"
// Output: "This is a sentence"
// Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.

// Input: s = "Myself2 Me1 I4 and3"
// Output: "Me Myself and I"
// Explanation: Sort the words in s to their original positions "Me1 Myself2 and3 I4", then remove the numbers.


func main() {
	s := "Myself2 Me1 I4 and3"
	fmt.Println(sortSentence(s))
}

func sortSentence(s string) string {
	
	mymap := map[int]string{}
	text := ""

	for _, v := range s {
		if(v != ' ') {
			if(v >= '0' && v <= '9') {
				mymap[int(v - 48)] = text
				text = ""
			}else {
				text = text + string(v)
			}
		}
	}

	text = "" 
	for i:= 1; i <= len(mymap); i++ {
		if(i == 1) {
			text = text + mymap[i]
		} else {
			text = text + " " + mymap[i]
		}
	}
	return text
}
