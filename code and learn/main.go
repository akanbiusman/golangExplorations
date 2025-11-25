package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 202
	var b uint8 = 141
	c := a &^ b
	fmt.Println(strconv.FormatUint(uint64(c), 2))
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(a & b)

}
