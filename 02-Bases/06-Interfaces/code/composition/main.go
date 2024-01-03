package main

type BookAttributes struct {
	// Title of the book
	Title string
	// Author of the book
	Author string
	// Year of publication
	Year int
}

type Book struct {
	// ID is the unique identifier of the book
	ID string
	// BookAttributes contains the attributes of the book
	BookAttributes
}