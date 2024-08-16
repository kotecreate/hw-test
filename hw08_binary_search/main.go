package main

import (
	"fmt"
)

func main() {
	slice := []int{5, 9, 14, 19, 21, 25, 35, 42}
	value := 19
	fmt.Printf("Массив: %v\n", slice)
	fmt.Printf("Найти: %v\n", value)
	fmt.Printf("Позиция числа: %v\n", binaryS(slice, value))
}

func binaryS(sl []int, val int) int {
	var valMid int
	var mid int
	left := 0
	right := (len(sl) - 1)
	for left <= right {
		mid = (left + right) / 2
		valMid = sl[mid]
		switch left <= right {
		case valMid == val:
			return mid
		case valMid < val:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}
