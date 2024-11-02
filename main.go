package main

import "fmt"

func sliceReverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)
	return slice
}

func findMax(array []int) int {
	max := array[0]
	for _, val := range array {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max value: ", max)
	return max
}
func main() {
	array := []int{10, 5, 2, 1, 34, 9, 23, 0}
	findMax(array)
	sliceReverse(array)
}
