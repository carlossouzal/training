package main

import "fmt"

func main() {
	var A, B, C, D int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	diference := (A*B - C*D)
	fmt.Printf("DIFERENCA = %d\n", diference)
}
