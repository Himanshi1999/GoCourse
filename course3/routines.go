package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func sortSubarray(subarray []int, ch chan []int) {
	fmt.Println("Sorting subarray:", subarray)
	sort.Ints(subarray)
	ch <- subarray
}

func routine() {
	fmt.Println("Enter a series of integers separated by spaces:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fields := strings.Fields(input)
	arr := make([]int, 0, len(fields))
	for _, f := range fields {
		num, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Invalid input:", f)
			return
		}
		arr = append(arr, num)
	}

	if len(arr) < 4 {
		fmt.Println("Please enter at least 4 integers.")
		return
	}

	n := len(arr)
	partSize := (n + 3) / 4

	ch := make(chan []int, 4)
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		if end > n {
			end = n
		}
		go sortSubarray(arr[start:end], ch)
	}

	sortedSubarrays := make([][]int, 0, 4)
	for i := 0; i < 4; i++ {
		sub := <-ch
		sortedSubarrays = append(sortedSubarrays, sub)
	}

	merged := merge(sortedSubarrays[0], sortedSubarrays[1])
	merged = merge(merged, sortedSubarrays[2])
	merged = merge(merged, sortedSubarrays[3])

	fmt.Println("Final sorted array:", merged)
}
