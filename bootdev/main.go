package main

import "fmt"

func main() {

}

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		for i := 3; i*i < n+1; i++ {

		}
	}
}
