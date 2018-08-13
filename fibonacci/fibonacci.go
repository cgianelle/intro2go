package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibo(n int) int {
	if n > 2 {
		return fibo(n-1) + fibo(n-2)
	}
	return 1
}

func main() {
	var startSeq, _ = strconv.Atoi(os.Args[1])
	var endSeq, _ = strconv.Atoi(os.Args[2])
	var seq []int
	fmt.Println(startSeq)
	fmt.Println(endSeq)

	for i := startSeq; i <= endSeq; i++ {
		seq = append(seq, fibo(i))
	}

	fmt.Println(seq)
}
