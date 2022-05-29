package main

import "fmt"

type Product struct {
	Codigo     int
	Quantidade int
	Valor      float32
}

func main() {
	peca1 := Product{}
	peca2 := Product{}

	fmt.Scanf("%d %d %f\n", &peca1.Codigo, &peca1.Quantidade, &peca1.Valor)
	fmt.Scanf("%d %d %f\n", &peca2.Codigo, &peca2.Quantidade, &peca2.Valor)

	total := (float32(peca1.Quantidade) * peca1.Valor) + (float32(peca2.Quantidade) * peca2.Valor)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
