package main

import (
	"fmt"
	"time"
)

var x int = 1

func task1() {
	fmt.Println("Task 1 start", x)
	x += 1
	fmt.Println("Task 1 end", x)
}

func task2() {
	fmt.Println("Task 2 start", x)
	x += 1
	fmt.Println("Task 2 end", x)
}

func race_condition() {
	go task1()
	go task2()
	time.Sleep(1 * time.Second)
	print(x)
}

/* Both goroutines access and modify the shared variable x.

The operation print is not atomic; it's actually made of three steps:

Read x value

Increment the value

Write the new value back

When both goroutines run concurrently:

They might read the same value at the same time, increment it, and overwrite each other's updates.

This leads to unpredictable results.*/
