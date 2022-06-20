package data

import "gorm.io/gorm"

const DBConnectionKey = "DBConnection"

type DbConnection interface {
	GetContext() (*gorm.DB, error)
	CloseConnection()
	Migrate(st ...interface{}) error
}
