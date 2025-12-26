package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkPrefix(s string) int {
	negative := strings.HasPrefix(s, "L")
	if negative {
		value, _ := strconv.Atoi(s[1:])
		return -value
	}
	value, _ := strconv.Atoi(s[1:])
	return value
}

func main() {
	dial := 50
	input := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	for _, i := range input {
		fmt.Printf("Current dial: %d\n", dial)
		fmt.Printf("Next input: %s\n", i)
		result := dial + checkPrefix(i)
		fmt.Printf("Result: %d\n", result)
		if result < 0 {
			dial = result + 100
		} else if result > 100 {
			dial = result - 100
		} else {
			dial = result
		}
		if dial == 100 {
			dial = 0
		}
		fmt.Println(dial)
	}
}
