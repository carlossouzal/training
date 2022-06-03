package main

import "fmt"

func main() {
	var howLong, velocity int

	fmt.Scan(&howLong)
	fmt.Scan(&velocity)

	total := float64(howLong*velocity) / 12.0
	fmt.Printf("%.3f\n", total)
}
