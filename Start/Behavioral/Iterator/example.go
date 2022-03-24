package main

import "fmt"

// Iterate using a callback function
func main() {
	// use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// Use IterateBooks to iterate via anonymous function
	lib.IterateBooks(func(b Book) error {
		fmt.Println("The book Author is: ", b.Author)
		return nil
	})

	// create a BookIterator
	bIterator := lib.createIterator()

	for bIterator.hasNext() {
		fmt.Println("Book is:", bIterator.next())
	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
