package main

import "diByWire/database"

type BookService struct {
	DB database.Database
}

func NewBookService(db database.Database) BookService {
	return BookService{
		DB: db,
	}
}

func (serv *BookService) GetBook() database.Book {
	return serv.DB.GetBook()
}
