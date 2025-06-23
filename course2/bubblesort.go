package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(sli []int, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}

func bubbleSort(sli []int) {
	n := len(sli)
	var swapped bool

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

}

func BubbleSort() {
	fmt.Println("Enter 10 numbers to sort.")
	fmt.Println("(Enter a number and press Enter. Repeat this 10 times.)")
	var numbers []int

	scanner := bufio.NewScanner(os.Stdin)

	for count := 0; count < 10; count++ {
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			number, _ := strconv.Atoi(input)
			numbers = append(numbers, number)
		} else {
			break
		}
	}

	bubbleSort(numbers)

	for _, num := range numbers {
		fmt.Print(num, " ")
	}

}
