package main

import "fmt"

func main() {
	printalphabet()
}

func printalphabet() {
	for r := 'a'; r <= 'z'; r++ {
		fmt.Print(string(r))
	}
}
