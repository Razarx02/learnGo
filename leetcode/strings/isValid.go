package main

import(
	"fmt"
)


// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.



func main() {
	s := "()[]{}"

	fmt.Println(isValid(s))
}

func isValid(s string) bool { 

	text := ""
	for _, value := range s {
		if (value == '(' || value == '{' || value == '[') {
			text = text + string(value)
		} else if (value == ')' &&  text != "" && rune(text[len(text) - 1]) == '('  ) {
			text = text[:len(text) - 1]	
		} else if (value == '}' && text != ""  &&  rune(text[len(text) - 1]) == '{'   ) {
			text = text[:len(text) - 1]	
		} else if (value == ']' &&  text != "" && rune(text[len(text) - 1]) == '['   ) {
			text = text[:len(text) - 1]	
		}  else {
			return false
		}	
	}
	if(len(text) == 0) {
		return true
	} else {
		return false
	}
	

}