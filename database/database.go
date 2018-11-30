package database

import (
	"github.com/whatvn/dqueue/helper"
	"github.com/whatvn/dqueue/models"
)

type Database interface {
	Init() error
}

func NewDatabase() Database {
	dbType := helper.GetDbType()
	switch dbType {
	case "mysql":
		return newMySqlDatabase()
	default:
		panic(message.NotImplementError)
	}
}
