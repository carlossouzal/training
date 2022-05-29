package main

import "fmt"

func volume(raio float64) float64 {
	pi := 3.14159
	return float64(4.0/3.0) * pi * (raio * raio * raio)
}

func main() {
	var raio float64
	fmt.Scan(&raio)

	fmt.Printf("VOLUME = %.3f\n", volume(raio))
}
