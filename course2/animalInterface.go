package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("Grass")
}

func (c Cow) Move() {
	fmt.Println("Walk")
}

func (c Cow) Speak() {
	fmt.Println("Moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("Worms")
}

func (b Bird) Move() {
	fmt.Println("Fly")
}

func (b Bird) Speak() {
	fmt.Println("Peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("Slither")
}

func (s Snake) Speak() {
	fmt.Println("Hsss")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animals := make(map[string]AnimalInterface)
	for {
		fmt.Println(">")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		request := strings.Fields(line)

		if len(request) != 3 {
			fmt.Println("Enter in the format of: newanimal/query name type/info")
			continue
		}

		command, name, param := request[0], request[1], request[2]

		switch command {
		case "newanimal":
			switch param {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type. Use cow, bird, or snake.")
				continue
			}
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query. Use eat, move, or speak.")
			}

		default:
			fmt.Println("Unknown command. Use newanimal or query.")
		}
	}

}
