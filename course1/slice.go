package main

import (
	"fmt"
	"sort"
	"strconv"
)

func slice() {
	slice := make([]int, 0, 3)
	var input string

	for {
		fmt.Println("Enter an integer (or 'X' to exit): ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		var value int
		value, _ = strconv.Atoi(input)
		slice = append(slice, value)

		sort.Ints(slice)
		fmt.Println("Sorted slice:", slice)
	}

}
