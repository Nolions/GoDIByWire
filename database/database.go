package database

type Database struct {
}

func NewDatabase() Database {
	return Database{}
}

type Book struct {
	Title string
	Price int
}

func (db *Database) GetBook() Book {
	return Book{
		Title: "HelloWork",
		Price: 0,
	}
}
