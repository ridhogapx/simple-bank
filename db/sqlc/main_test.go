package db

import (
	"database/sql"
	"fmt"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		fmt.Printf("Failed to connect database: %v", m)
	}
	testQueries = New(conn)
}
