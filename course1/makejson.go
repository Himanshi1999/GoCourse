package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func makejson() {
	type People struct {
		Name    string
		Address string
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	fmt.Print("Enter the address: ")
	addressInput, _ := reader.ReadString('\n')
	address := strings.TrimSpace(addressInput)

	person := People{
		Name:    name,
		Address: address,
	}

	b, _ := json.Marshal(person)

	os.Stdout.Write(b)

}
