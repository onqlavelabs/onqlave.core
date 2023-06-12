package database

import "gorm.io/gorm"

const DBConnectionKey = "DBConnection"

type Connection interface {
	GetContext() (*gorm.DB, error)
	CloseConnection()
	Migrate(schemas ...interface{}) error
	Seed() error // we need to implement this properly!
}
