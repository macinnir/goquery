package testassets

import (
	"context"
	"database/sql"
)

// MockDB is a mock implementation of the IDB interface for testing purposes
type MockDB struct {
	// You can add fields here to track calls and parameters if needed

}

func (db *MockDB) Close() {
}

func (db *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (db *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return nil, nil
}

func (db *MockDB) ExecMany(stmts []string, chunkSize int) error {
	return nil
}

func (db *MockDB) Host() string {
	return "localhost"
}

func (db *MockDB) Name() string {
	return "testdb"
}

func (db *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (db *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}
