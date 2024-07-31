package main

import (
	"fmt"
	"log"
)

func main() {
	var arr = make([]int, 3)

	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])

		if err != nil {
			log.Fatal(err)
		}
	}

	sorted := bubble(arr)
	fmt.Print(sorted[1])
}

func bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max_found := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				max_found = true
			}
		}

		if max_found == false {
			break
		}
	}

	return arr
}
