package repository

import "testing"

func TestNewDB(t *testing.T) {
	db, err := LoadSqlConnection()
	if err != nil {
		t.Fatalf("Failed connection to MySQL")
	}
	defer db.Close()
}
