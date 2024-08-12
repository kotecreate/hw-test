package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStCompCaseTable(t *testing.T) {
	testCases := []struct {
		name   string
		b1, b2 Book
		expect bool
	}{
		{
			name: "testSC",
			b1: Book{
				id:     1,
				title:  "1984",
				author: "George Orwell",
				year:   1949,
				size:   196,
				rate:   4.6,
			},
			b2: Book{
				id:     2,
				title:  "Le Petit Prince",
				author: "Antoine Marie Jean-Baptiste Roger vicomte de Saint-Exupery",
				year:   1943,
				size:   128,
				rate:   4.8,
			},
			expect: true,
		},
		{
			name: "testSC",
			b1: Book{
				id:     3,
				title:  "Зеленая миля",
				author: "Стивен Кинг",
				year:   1997,
				size:   384,
				rate:   4.7,
			},
			b2: Book{
				id:     4,
				title:  "Граф Монте-Кристо",
				author: "Александр Дюма",
				year:   2021,
				size:   984,
				rate:   4.8,
			},
			expect: false,
		},
		{
			name: "testSC",
			b1: Book{
				id:     5,
				title:  "Властелин колец",
				author: "Джон Р.Р.Толкин",
				year:   2002,
				size:   992,
				rate:   4.8,
			},
			b2: Book{
				id:     6,
				title:  "Десять негритят",
				author: "Агата Кристи",
				year:   1939,
				size:   256,
				rate:   4.6,
			},
			expect: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := SC(tC.b1, tC.b2)
			assert.Equal(t, tC.expect, actual)
		})
	}
}
