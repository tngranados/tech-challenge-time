package storage

import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// DBSession is an interface that both `sql.DB` and `sql.Tx` implements.
type DBSession interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// TxFunc is a function used to execute statements and queries against a
// database.
type TxFunc func(DBSession) error

// Transaction creates a new transaction and handles the rollback and commit
// logic based on the error object returned by the `Transaction`
func Transaction(db *sql.DB, fn TxFunc) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback and panic again.
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// An error occurred, rollback.
			_ = tx.Rollback()
		} else {
			// No errors, commit.
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

// NewDatabase returns the DB connection.
func NewDatabase(databasePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON", nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewTestDB creates a new database suited to be used for tests. The cleanup
// function is used to remove the temporary database.
func NewTestDB(tb testing.TB) (db *sql.DB, cleanup func()) {
	dir, err := ioutil.TempDir("", "test_db_")
	if err != nil {
		tb.Fatal(err)
	}

	tmpFile, err := ioutil.TempFile(dir, "")
	if err != nil {
		tb.Fatal(err)
	}

	db, err = NewDatabase(tmpFile.Name())
	if err != nil {
		tb.Fatal(err)
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			tb.Fatal(err)
		}

		err = os.RemoveAll(dir)
		if err != nil {
			tb.Fatal(err)
		}
	}
}
