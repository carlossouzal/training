// As vezes o input simplementes pula a segunda entrada
// Parece ser alguma coisa com o \n no buffer
// Porém não resolvi de uma forma definitiva, só acontece as vezes

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func distance(P1 Point, P2 Point) float64 {
	distance := math.Pow((P2.x-P1.x), 2) + math.Pow((P2.y-P1.y), 2)
	return math.Sqrt(distance)
}

func main() {
	var P1, P2 Point

	fmt.Scanf("%f %f\n", &P1.x, &P1.y)
	fmt.Scanf("%f %f\n", &P2.x, &P2.y)

	fmt.Printf("%.4f\n", distance(P1, P2))
}
