package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
)

func LoadSqlConnection() (*sql.DB, error) {
	var err error
	var sqlConnection *sql.DB

	once.Do(func() {
		driverName := "mysql"

		url := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			"root",
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			"3306",
			"advisories",
		)
		sqlConnection, err = sql.Open(driverName, url)
		if err != nil {
			return
		}
		err = sqlConnection.Ping()
	})
	return sqlConnection, err
}

func NewDB() *sql.DB {
	db, err := LoadSqlConnection()
	if err != nil {
		panic(err)
	}
	return db
}
