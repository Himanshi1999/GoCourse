package main

import (
	"fmt"
)

func trunc() {
	var floatNumber float64
	fmt.Println("Enter a floating point number")
	fmt.Scan(&floatNumber)
	fmt.Printf("Int value of your floating point number is %d\n", int(floatNumber))

	fmt.Println("Enter a floating point number again")
	fmt.Scan(&floatNumber)
	fmt.Printf("Int value of your floating point number is %d\n", int(floatNumber))
}
