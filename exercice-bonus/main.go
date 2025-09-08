package main

import "fmt"

func main() {
	fmt.Println(BinaryToDecimal("1011"))
	fmt.Println(BinaryToDecimal("1111"))
	fmt.Println(BinaryToDecimal("1000"))
}

func BinaryToDecimal(s string) int {
	if s == "" {
		return 0
	}

	last := s[len(s)-1]
	digit := 0

	switch last {
	case '0':
		digit = 0
	case '1':
		digit = 1
	}

	return BinaryToDecimal(s[:len(s)-1])*2 + digit
}
