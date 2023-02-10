package main

import (
	"diByWire/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

var book database.Book

func init() {
	book = database.Book{
		Title: "HelloWork",
		Price: 0,
	}
}

func TestGetBook(t *testing.T) {
	serv := NewBookService()

	assert.Equal(t, book, serv.GetBook())
}
