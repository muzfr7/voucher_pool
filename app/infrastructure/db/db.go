package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Database defines contracts for database connection.
type Database interface {
	GetConnection() (*gorm.DB, error)
}

// database implements Database interface.
type database struct {
	driver string
	user   string
	pass   string
	host   string
	port   string
	name   string
}

// NewDatabase returns a new instance of the database.
func NewDatabase(driver, user, pass, host, port, name string) Database {
	return &database{
		driver: driver,
		user:   user,
		pass:   pass,
		host:   host,
		port:   port,
		name:   name,
	}
}

// GetConnection will initialize and return a new db connection.
func (d *database) GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", d.user, d.pass, d.host, d.port, d.name)
	db, err := gorm.Open(d.driver, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
