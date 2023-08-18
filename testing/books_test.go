package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReadBooks(t *testing.T) {
	csvFile, err := os.Open("../data/books.csv")
	if err != nil {
		t.Fatal("Can't open CSV file")
	}
	expectedBooksLen := 3
	books, err := ReadBooks(csvFile)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		t.Fatal("Can't read CSV data")
	}
	actualBooksLen := len(books)

	if expectedBooksLen != actualBooksLen {
		t.Errorf("Unexpected number of books, got: %d, want: %d.", actualBooksLen, expectedBooksLen)
	}

	expectedBooks := []Book{
		{
			title:  "The Guardians",
			author: "John Grisham",
			year:   2019,
		},
		{
			title:  "To Kill a Mockingbird",
			author: "Harper Lee",
			year:   2005,
		},
		{
			title:  "War and Peace",
			author: "Leo Tolstoy",
			year:   1982,
		},
	}

	for i, b := range books {
		if expectedBooks[i].title != b.title {
			t.Errorf("Unexpected title, got: %s, want: %s.", b.title, expectedBooks[i].title)
		}
		if expectedBooks[i].author != b.author {
			t.Errorf("Unexpected author, got: %s, want: %s.", b.author, expectedBooks[i].author)
		}
		if expectedBooks[i].year != b.year {
			t.Errorf("Unexpected year, got: %d, want: %d.", b.year, expectedBooks[i].year)
		}
	}
}
