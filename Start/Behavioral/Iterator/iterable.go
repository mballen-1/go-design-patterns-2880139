package main

// The IterableCollection interface defines the createIterator
// function, which returns an iterator object
type IterableCollection interface {
	createIterator() iterator
}

// The iterator interface contains the hasNext and next functions
// which allow the collection to return items as needed
type iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	books   []Book
	current int
}

func (b *BookIterator) hasNext() bool {
	return b.current <= len(b.books)-1
}

func (b *BookIterator) next() *Book {
	if b.hasNext() {
		old := b.books[b.current]
		b.current++
		return &old
	}
	return nil
}
