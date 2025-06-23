package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func trimTo20(s string) string {
	if len(s) > 20 {
		return s[:20]
	}
	return s
}

func main() {
	var fileName string
	fmt.Print("Enter the name of text file: ")
	fmt.Scan(&fileName)

	var people []Name

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)

		if len(parts) == 2 {
			person := Name{
				Fname: trimTo20(parts[0]),
				Lname: trimTo20(parts[1]),
			}
			people = append(people, person)
		} else {
			fmt.Println("Not able to parse line", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, p := range people {
		fmt.Printf("First name is %s and last name is %s\n", p.Fname, p.Lname)
	}
}
