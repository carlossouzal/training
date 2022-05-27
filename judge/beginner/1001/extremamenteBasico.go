package main

import (
	"fmt"
)

func main() {
	var A,
		B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	x := A + B
	fmt.Println("X =", x)
}
