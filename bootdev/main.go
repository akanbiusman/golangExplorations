package main

import "fmt"

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	x := 100
	for idx := range x {
		if idx%3 == 0 {
			fmt.Printf("num: %v, fizz\n", idx)
		} else if idx%5 == 0 {
			fmt.Printf("num: %v, buzz\n", idx)
		} else if idx%3 == 0 && idx%5 == 0 {
			fmt.Printf("num: %v, fizzbuzz\n", idx)
		} else {
			continue
		}
	}
}
