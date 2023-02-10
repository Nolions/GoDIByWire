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

func TestGetBookByManual(t *testing.T) {
	db := database.NewDatabase()
	serv := NewBookService(db)

	assert.Equal(t, book, serv.GetBook())
}

func TestGetBookByWire(t *testing.T) {
	serv := InitBookService()

	assert.Equal(t, book, serv.GetBook())
}
