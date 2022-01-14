package main

import(
	"fmt"
)


// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).

// The matching should cover the entire input string (not partial).

func main() {
	
	s :=  "aab"
	p := "a*b"

	fmt.Println(isMatch(s,p))

}


func isMatch(s string, p string) bool {
	
	index_pattern := 0
	result := false


	for index, value := range s {

		if(value != '*') {
			if(p[index_pattern] == '*') {
				
			}
		}
	}

	

	return result
}