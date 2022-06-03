package main

import "fmt"

func main() {
	var seconds int

	fmt.Scan(&seconds)

	hours, minutes, seconds := parseTime(seconds)
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}

func parseTime(seconds int) (int, int, int) {
	remainderSeconds := seconds % 60
	minutes := seconds / 60
	hours := minutes / 60
	minutes = minutes % 60

	return hours, minutes, remainderSeconds
}
