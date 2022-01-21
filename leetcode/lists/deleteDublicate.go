package main

import (
	"fmt"
)



type ListNode struct {
	Val int
	Next *ListNode 
}



func main() {

	// array := []int{}
	// list := &ListNode{}

	// for _, v := range array {
	// 	addItemToLists(list, v)	
	// }
	// removeNoneItem(list)
	fmt.Println(deleteDuplicates(nil))
	// removeNoneItem(result)
	// fmt.Println(result)
	// printLists(*result)
}


func deleteDuplicates(head *ListNode) *ListNode { 
	if (head != nil) {
	result := &ListNode{}
	current := head
	
		
	
	data := current.Val



	// fmt.Println("aaa")	
	addItemToLists(result, data)
	for ( current != nil) {

		if ( data != current.Val) {
			data = current.Val
			addItemToLists(result, data)
		}
		current = current.Next
	}

	removeNoneItem(result)
	return result
}
	return nil
}


//  add Item to Lists
func addItemToLists(List *ListNode, item int) {
	current := List
	newItem := &ListNode{Val:item, Next:nil}
	for (current.Next != nil) {
		current = current.Next
	}
	current.Next = newItem
}

// remove None item FIRST
func removeNoneItem(List *ListNode) {
	current := List
	current = current.Next
	*List = *current
}



// Function to print Linked list
func printLists(list ListNode){
	
	for (list.Next != nil) {
		fmt.Println(list.Val)
		list = *list.Next
	}	
	fmt.Println(list.Val)
	
}