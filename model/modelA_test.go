package model

import (
	"testing"

	"test-scrutinizer/log"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertNewRecords(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("could not mock SQL service : %s", err)
	}
	defer db.Close()
	Name := "ab"
	City := "def"

	mock.ExpectExec("INSERT INTO new_table").WithArgs(Name, City).WillReturnResult(sqlmock.NewResult(1, 1))
	myDB := NewDummyStructure(db)
	myDB.InsertNewRecords("abc", "def")

}
