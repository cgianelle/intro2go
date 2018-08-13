package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func noNegativeNumbers() {
	fmt.Println("NO NEGATIVE NUMBERS PLEASE")
}

func nothingGreaterThan100() {
	fmt.Println("Please, no mumbers greater than 100")
}

func main() {
	var f, _ = strconv.Atoi(os.Args[1])
	var factorial big.Int

	if f < 0 {
		defer noNegativeNumbers()
	} else if f > 100 {
		defer nothingGreaterThan100()
	} else {
		factorial.MulRange(1, int64(f))
		fmt.Println(&factorial)
	}
}
