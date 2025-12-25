package main

import (
	"fmt"
	"strings"
)

func checkPrefix(s string) string {
	negative := strings.HasPrefix(s, "L")
	if negative {
		return "NEGATIVE"
	}
	return "POSITIVE"
}

func main() {
	// startingPoint := 50
	input := []string{"L68", "L30", "R48"}

	for _, i := range input {
		fmt.Println(checkPrefix(i))
	}
}
