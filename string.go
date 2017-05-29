package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "abc中国"
	fmt.Println(s1)

	for i := 0; i < len(s1); i++ {
		fmt.Println(string(s1[i]), "\t", reflect.TypeOf(s1[i]))
	}
	fmt.Println()

	for index, data := range s1 {
		fmt.Println(index, string(data), reflect.TypeOf(data))
	}

	s2 := "abcdeft1234"

	b1 := []byte(s2)
	b1[0] = 'x'

	fmt.Println(string(b1))
}
