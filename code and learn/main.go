package main

import "fmt"

func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		classes   []string
	}

	foo := student{
		firstName: "usman",
		lastName:  "akanbi",
		age:       10,
	}

	fmt.Println(foo.lastName)
}
