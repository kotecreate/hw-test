package main

import (
	"encoding/json"
	"fmt"

	bookpb "github.com/kotecreate/hw-test/hw09_serialize/pb"
	"google.golang.org/protobuf/proto"
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
	MarshalJSON() ([]byte, error)
}
type Unmarshaller interface {
	UnmarshalJSON([]byte) error
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Book) UnmarshalJSON(data []byte) (*Book, error) {
	err := json.Unmarshal(data, &b)
	if err != nil {
		fmt.Printf("Не удалось из-за: %v\n", err)
		return nil, err
	}
	return b, nil
}

func main() {
	var err error
	var mybook Book
	bookList := make(map[int]Book)
	bookList[1] = Book{
		ID:     1,
		Title:  "Идиот",
		Author: "Л. Н. Толстой",
		Year:   1869,
		Size:   135,
		Rate:   4.23,
	}
	bookList[2] = Book{
		ID:     2,
		Title:  "Преступление и наказание",
		Author: "Ф. Достоевский",
		Year:   1866,
		Size:   150,
		Rate:   4.12,
	}

	mybookProto := bookpb.BookList{
		Books: []*bookpb.BookInfo{
			{
				Id:     3,
				Title:  "Go Programming",
				Author: "John Doe",
				Year:   2020,
				Size:   350,
				Rate:   4.5,
			},
			{
				Id:     4,
				Title:  "1984",
				Author: "George Orwell",
				Year:   2024,
				Size:   320,
				Rate:   4.8,
			},
		},
	}
	dataM, err := serialize(bookList)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}
	mybook = bookList[1]

	mybookSerialized, err := serialize(mybook)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}

	fmt.Println(dataM)
	fmt.Println(mybookSerialized)

	newBookDecode, err := mybook.UnmarshalJSON(dataM)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}
	fmt.Printf("Book: %+v\n", newBookDecode)

	newBook, err := deserialize(mybookSerialized)
	if err != nil {
		fmt.Printf("Не удалось из-за: %+v\n", err)
	}

	fmt.Printf("Book: %+v\n", newBook)
	fmt.Printf("Title: %s\n", newBook.Title)
	fmt.Printf("Author: %s\n", newBook.Author)

	buf, _ := proto.Marshal(&mybookProto)
	fmt.Println("Proto: ", buf)
}

func serialize(b any) ([]byte, error) {
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
