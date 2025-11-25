package main

import "fmt"

func main() {
	var a = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(a[1][1])

	a[1][1] = 5
	fmt.Println(a[1][1])
}
