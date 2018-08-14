package main

import (
	"fmt"
)

func main() {

	operations := map[string][3]int{
		"a": [3]int{4, 5, 0},
		"d": [3]int{0, 3, 0},
		"m": [3]int{4, 8, 0},
		"s": [3]int{1, 4, 0},
	}

	functions := map[string]func(int, int) int{
		"a": func(a int, b int) int { return a + b },
		"s": func(a int, b int) int { return a - b },
		"m": func(a int, b int) int { return a * b },
		"d": func(a int, b int) int { return a / b },
	}

	for op, values := range operations {
		operation := functions[op]
		values[2] = operation(values[0], values[1])
		operations[op] = values
	}

	fmt.Println(operations)
}
