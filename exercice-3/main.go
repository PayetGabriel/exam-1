package main

import "fmt"

func main() {
	printdigits()
}

func printdigits() {
	for r := 48; r <= 57; r++ {
		fmt.Print(string(r))
	}
	fmt.Print("\n")
}
