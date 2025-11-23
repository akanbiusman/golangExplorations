package main

import (
	"errors"
	"fmt"
)

type student struct {
	firstName string
	lastName  string
}

func main() {
	newStudent := student{
		firstName: "usman",
		lastName:  "akanbi",
	}

	prev, err := updateLastName(&newStudent, "akanbi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prev)
}

func updateLastName(s *student, newlastName string) (*string, error) {
	if newlastName == "" {
		return nil, errors.New("empty last name")
	}
	previous := s.lastName
	s.lastName = newlastName
	return &previous, nil
}
