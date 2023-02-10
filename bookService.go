package main

import "diByWire/database"

type BookService struct {
}

func NewBookService() BookService {
	return BookService{}
}

func (serv *BookService) GetBook() database.Book {
	db := database.NewDatabase()
	return db.GetBook()
}
