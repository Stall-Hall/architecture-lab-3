package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "lab3",
		User:       "postgres",
		Password:   "qwe",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:qwe@localhost/lab3?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
