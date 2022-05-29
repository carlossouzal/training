package main

import (
	"fmt"
	"math"
)

func bigger(a int, b int) int {
	return (a + b + int(math.Abs(float64(a-b)))) / 2
}

func main() {
	var A, B, C int

	fmt.Scanf("%d %d %d\n", &A, &B, &C)

	maior := bigger(A, B)
	maior = bigger(maior, C)

	fmt.Printf("%d eh o maior\n", maior)
}
