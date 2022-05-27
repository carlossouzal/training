package main

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	soma := A + B

	fmt.Printf("SOMA = %d\n", soma)
}
