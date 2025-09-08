package main

import "fmt"

func main() {
	fmt.Println(pgcd(18, 15))
	fmt.Println(pgcd(6, 12))
	fmt.Println(pgcd(24, 10))
}

func pgcd(nb1 int, nb2 int) int {
	if nb1 < nb2 {
		for i := nb1; i >= 0; i-- {
			if nb1%i == 0 {
				if nb2%i == 0 {
					return i
				}
			}
		}
	} else if nb1 > nb2 {
		for i := nb2; i >= 0; i-- {
			if nb2%i == 0 {
				if nb1%i == 0 {
					return i
				}
			}
		}
	}
	return 0
}
