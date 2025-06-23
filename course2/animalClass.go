package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func animalClass() {

	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		request := strings.SplitN(line, " ", 2)

		if len(request) != 2 {
			fmt.Println("Enter in the format of: animal action")
			continue
		}
		animalName := strings.ToLower(request[0])
		action := strings.ToLower(request[1])

		animal, isAvailable := animals[animalName]

		if !isAvailable {
			fmt.Println("Please check animal name, available animals are cow, bird, snake")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "speak":
			animal.Speak()
		case "move":
			animal.Move()
		}
	}

}
