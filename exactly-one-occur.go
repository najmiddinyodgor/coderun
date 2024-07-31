package main

import (
	"fmt"
	"log"
)

func main() {
	var count int

	_, err := fmt.Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	var arr = make([]int, count)

	for i := 0; i < count; i++ {
		_, err := fmt.Scan(&arr[i])

		if err != nil {
			log.Fatal(err)
		}
	}

	var hashmap = make(map[int]int)

	for i := 0; i < count; i++ {
		_, ok := hashmap[arr[i]]
		if ok {
			hashmap[arr[i]] = 0
		} else {
			hashmap[arr[i]] = 1
		}
	}

	result := 0

	for _, number := range hashmap {
		result += number
	}

	fmt.Print(result)
}
