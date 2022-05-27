package main

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	product := A * B

	fmt.Printf("PROD = %d\n", product)
}
