package main

import "fmt"

func main() {
	a := squareNum(5)
	fmt.Println(a)

	b := make([]int, 5)
	b[1] = 10

	fmt.Println(b)

	c := appendSlice([]int{1, 2, 3})
	fmt.Println(c)
}

func squareNum(num int) int {
	return num * num
}

func appendSlice(x []int) []int {
	x = append(x, 11)
	return x
}
