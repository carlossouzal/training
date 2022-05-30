package main

import "fmt"

func main() {
	var distance int
	var consumption float64

	fmt.Scan(&distance)
	fmt.Scan(&consumption)

	media := float64(distance) / consumption
	fmt.Printf("%.3f km/l\n", media)
}
