package driver

import (
	"database/sql"
	"fmt"

	"test-scrutinizer/config"
)

// DbConn : function for database connection
func DbConn(localConfig config.Provider) (*sql.DB, error) {
	dbUser := localConfig.GetString("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("connection cannot be made as field for dbUser is not defined")
	}
	dbPass := localConfig.GetString("DB_PASSWORD")
	if dbPass == "" {
		return nil, fmt.Errorf("connection cannot be made as field for dbPass is not defined")
	}
	dbName := localConfig.GetString("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("connection cannot be made as field for dbName is not defined")
	}
	dbIP := localConfig.GetString("DB_HOST")
	if dbIP == "" {
		return nil, fmt.Errorf("connection cannot be made as field for dbIP is not defined")
	}
	dbPort := localConfig.GetString("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("connection cannot be made as field for dbPort is not defined")
	}
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		return nil, fmt.Errorf("connection cannot be made with mysql : %s", err.Error())
	}
	return db, nil
}
