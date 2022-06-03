package main

import "fmt"

func main() {
	var daysOld int

	fmt.Scan(&daysOld)
	years, month, days := parse(daysOld)

	fmt.Printf("%d ano(s)\n", years)
	fmt.Printf("%d mes(es)\n", month)
	fmt.Printf("%d dia(s)\n", days)
}

func parse(days int) (int, int, int) {
	years := days / 365
	month := (days % 365) / 30
	remaindersDays := (days % 365) % 30

	return years, month, remaindersDays
}
