package main

import (
	"fmt"
	"time"
)

type Book struct {
	title    string
	author   string
	pages    int
	flips    int
	isSaved  bool
	isRented bool
	readTime time.Time
}

func main() {
	gow := Book{
		title:    "god of war",
		author:   "usman",
		pages:    10,
		flips:    3,
		isSaved:  false,
		isRented: false,
		readTime: time.Now(),
	}

	saveBook(&gow)
	fmt.Println(gow.isSaved)
	fmt.Println(gow.readTime)
	flipBook(&gow)
}

func flipBook(book *Book) {
	for i := 0; i < book.pages; i++ {
		if i == book.flips {
			pagesLeft := book.pages - book.flips
			fmt.Printf("You have %v pages left to read", pagesLeft)
		}
	}
}

func saveBook(book *Book) {
	book.isSaved = true
	book.readTime = time.Now()
}
