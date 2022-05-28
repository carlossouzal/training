package main

import "fmt"

func main() {
	var id, hora int
	var precoHora float32

	fmt.Scan(&id)
	fmt.Scan(&hora)
	fmt.Scan(&precoHora)

	salario := float32(hora) * precoHora
	fmt.Printf("NUMBER = %d\n", id)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
