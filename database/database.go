package database

import (
	"github.com/whatvn/dqueue/helper"
	"errors"
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
		panic(errors.New("not implement"))
	}
}
