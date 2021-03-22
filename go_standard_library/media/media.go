package media

import "strings"

type Catelogable interface {
	NewMovie(title string, rating Rating, boxOffice float32)
	GetTitle() string
	GetRating() Rating
	GetBoxOffice() float32
	SetTitle(newTitle string)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R"
	G    = "G"
	PG   = "PG"
	PG13 = "PG-13"
	NC17 = "NC-17"
)

func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

func (m *Movie) GetRating() Rating {
	return m.rating
}

func (m *Movie) SetTitle(newTitle string) {
	m.title = newTitle
}
