package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Print(nextInt())
	fmt.Print(nextInt())
	fmt.Print(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}