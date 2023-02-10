// wireBookService.go
//go:build wireinject
// +build wireinject

package main

// generator code by command wire

import (
	"diByWire/database"
	"github.com/google/wire"
)

func InitBookService() BookService {
	panic(wire.Build(NewBookService, database.NewDatabase))
}
