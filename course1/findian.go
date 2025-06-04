package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian() {
	var inputString string
	fmt.Println("Enter input string")

	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	inputString = strings.ToLower(strings.TrimSpace(inputString))

	if strings.HasPrefix(inputString, "i") && strings.HasSuffix(inputString, "n") && strings.Contains(inputString, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
