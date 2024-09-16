package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float32 `json:"rate"`
}

type Marshaller interface {
	marshalJSON() ([]byte, error)
}
type Unmarshaller interface {
	unmarshalJSON([]byte) error
}

func (b *Book) marshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Book) unmarshalJSON(data []byte) (*Book, error) {
	err := json.Unmarshal(data, &b)
	if err != nil {
		fmt.Printf("Не удалось из-за: %v\n", err)
		return nil, err
	}
	return b, nil
}

func main() {
	var err error
	mybook := &Book{
		ID:     12,
		Title:  "Gone Girl",
		Author: "Marcel Proust",
		Year:   1956,
		Size:   160,
		Rate:   4.2,
	}

	dataM, err := serialize(mybook)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}

	fmt.Println(dataM)

	newBookDecode, err := mybook.unmarshalJSON(dataM)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}
	fmt.Printf("Book: %+v\n", newBookDecode)

	newBook, err := deserialize(dataM)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}

	fmt.Printf("Book: %+v\n", newBook)
	fmt.Printf("Title: %s\n", newBook.Title)
	fmt.Printf("Author: %s\n", newBook.Author)
}

func serialize(b *Book) ([]byte, error) {
	jsonData, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func deserialize(jsonData []byte) (*Book, error) {
	var book Book
	err := json.Unmarshal(jsonData, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
