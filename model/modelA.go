package model

import (
	"database/sql"

	"test-scrutinizer/log"

	_ "github.com/go-sql-driver/mysql" // MYSQL
)

// Dummy : interface for Dummy
type Dummy interface {
	InsertNewRecords(string, string)
}

// DummyStructure : structure for scrutinizer_test database connection
type DummyStructure struct {
	DB *sql.DB
}

// InsertNewRecords :
func (ds *DummyStructure) InsertNewRecords(nam string, cit string) {
	query := "INSERT INTO new_table(name, city)VALUES(?, ?)"
	_, err := ds.DB.Exec(query, nam, cit)
	if err != nil {
		log.Fatalf("could not insert into new_table : %s", err)
	}
}

// NewDummyStructure : creates receiver for DummyStructure
func NewDummyStructure(db *sql.DB) Dummy {
	return &DummyStructure{
		DB: db,
	}
}
