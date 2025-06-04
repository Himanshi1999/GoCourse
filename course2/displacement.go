package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration, initial_velocity, initial_displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5*acceleration)*(time*time) + (initial_velocity * time) + initial_displacement
	}
}

func main() {
	var acceleration, initial_velocity, initial_displacement, time float64

	fmt.Print("Enter the value for acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Print("\nEnter the value for initial velocity: ")
	fmt.Scan(&initial_velocity)

	fmt.Print("\nEnter the value for initial displacement: ")
	fmt.Scan(&initial_displacement)

	fmt.Print("\nEnter the value for initial time: ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)

	fmt.Printf("Acceleration after time %f is %f", time, fn(time))

}
