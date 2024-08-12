package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

type Comparison struct {
	field int
}

func NewFi(field int) *Comparison {
	return &Comparison{
		field: field,
	}
}

func (c *Comparison) Compare(book1, book2 *Book) bool {
	switch c.field {
	case 0:
		return book1.year > book2.year
	case 1:
		return book1.size > book2.size
	case 2:
		return book1.rate > book2.rate
	default:
		return false
	}
}

const (
	Year = iota
	Size
	Rate
)

// type Stringer interface {
// 	String() string
// }

func (b *Book) printAll() {
	fmt.Println("Идентификационный номер:", b.id)
	fmt.Println("Название книги:", b.title)
	fmt.Println("Автор книги:", b.author)
	fmt.Println("Год выпуска книги:", b.year)
	fmt.Println("Количество страниц:", b.size)
	fmt.Println("Рейтинг книги:", b.rate)
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float64 {
	return b.rate
}

func (b *Book) SetID(newID int) {
	b.id = newID
}

func (b *Book) SetTitle(newTitle string) {
	b.title = newTitle
}

func (b *Book) SetAuthor(newAuthor string) {
	b.author = newAuthor
}

func (b *Book) SetYear(newYear int) {
	b.year = newYear
}

func (b *Book) SetSize(newSize int) {
	b.size = newSize
}

func (b *Book) SetRate(newRate float64) {
	b.rate = newRate
}

func main() {
	book1 := Book{
		id:     1,
		title:  "1984",
		author: "George Orwell",
		year:   1949,
		size:   196,
		rate:   4.6,
	}
	book2 := Book{
		id:     2,
		title:  "Le Petit Prince",
		author: "Antoine Marie Jean-Baptiste Roger vicomte de Saint-Exupery",
		year:   1943,
		size:   128,
		rate:   4.8,
	}
	book2.printAll()
	book2.SetRate(4.9)
	book2.printAll()
	SC(book1, book2)
}

func SC(b1, b2 Book) bool {
	book1, book2 := b1, b2

	comp := NewFi(Year)
	fmt.Printf("Сравнение Книги 1 и Книги 2 по Году: %v\n", comp.Compare(&book1, &book2))
	return comp.Compare(&book1, &book2)
}
