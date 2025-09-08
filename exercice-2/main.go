package main

import "fmt"

func main() {
	fmt.Println(CountLetters("Hello World !"))
	fmt.Println(CountLetters("125@6 a"))
	fmt.Println(CountLetters("exam 1 golang"))
}

func CountLetters(s string) int {
	result := 0
	for _, r := range s {
		if (r >= 97 && r <= 122) || (r >= 65 && r <= 90) {
			result++
		}
	}
	return result
}
