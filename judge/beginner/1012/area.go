// Note:
// Ao usar float32 a precisão menor do float acaba gerando um erro em 5% dos casos.
// Séria possivel resolver usando o math.Round() para floats.
// Porém o beecrowd apresenta um "compilation error", neste caso a melhor forma de resolver é alterando de float32 para float64

package main

import "fmt"

func areaTriangulo(base float64, altura float64) float64 {
	return base * altura / 2
}

func areaCirculo(raio float64) float64 {
	pi := float64(3.14159)
	return pi * (raio * raio)
}

func areaTrapezio(baseMaior float64, baseMenor float64, altura float64) float64 {
	return (baseMaior + baseMenor) * altura / 2
}

func areaRetangulo(ladoA float64, ladoB float64) float64 {
	return ladoA * ladoB
}

func main() {
	var A, B, C float64
	fmt.Scanf("%f %f %f\n", &A, &B, &C)

	fmt.Printf("TRIANGULO: %.3f\n", areaTriangulo(A, C))
	fmt.Printf("CIRCULO: %.3f\n", areaCirculo(C))
	fmt.Printf("TRAPEZIO: %.3f\n", areaTrapezio(A, B, C))
	fmt.Printf("QUADRADO: %.3f\n", areaRetangulo(B, B))
	fmt.Printf("RETANGULO: %.3f\n", areaRetangulo(A, B))
}
