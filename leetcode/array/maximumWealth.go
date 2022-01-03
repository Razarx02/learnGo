package main

import (
	"fmt"
)


func main() {
	array := [][]int{{1,5},{7,3},{3,5}}

	value := maximumWealth(array)

	fmt.Println(value)

}

func maximumWealth(accounts [][]int) int {

	result := 0

	for _, v := range accounts {
		allValue := 0
		for _, v_in := range v {
			allValue = allValue + v_in
		} 	

		if (allValue >= result) {
			result = allValue
		}
	}

	return result
}