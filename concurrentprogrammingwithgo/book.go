package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Book 1",
		Author:        "Author 1",
		YearPublished: 1,
	},
	Book{
		ID:            2,
		Title:         "Book 2",
		Author:        "Author 2",
		YearPublished: 2,
	},
	Book{
		ID:            3,
		Title:         "Book 3",
		Author:        "Authro 3",
		YearPublished: 3,
	},
	Book{
		ID:            4,
		Title:         "Book 4",
		Author:        "Authro 4",
		YearPublished: 4,
	},
}
