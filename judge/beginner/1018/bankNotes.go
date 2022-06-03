package main

import "fmt"

func main() {
	bankNotes := []int{100, 50, 20, 10, 5, 2, 1}
	var value int

	fmt.Scan(&value)
	result := decompose(value, bankNotes)

	fmt.Println(value)
	for i := 0; i < len(bankNotes); i++ {
		fmt.Printf("%d nota(s) de R$ %d,00\n", result[i], bankNotes[i])
	}
}

func decompose(value int, notes []int) []int {
	response := make([]int, len(notes))

	for amount, i := 0, 0; i < len(notes) && amount < value; i++ {
		quantity := (value - amount) / notes[i]
		response[i] = quantity
		amount += quantity * notes[i]
	}

	return response
}
