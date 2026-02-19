package main

import (
	"book/book"
	"fmt"
)

func main() {
	var book book.Book

	book.Title = "A clockwork Orange"

	fmt.Printf("The book's title is: \"%s\"\n", book.Title)
}
