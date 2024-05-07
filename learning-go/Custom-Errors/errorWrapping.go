package main

import (
	"database/sql"
	"fmt"
)

func queryAndProcess() {
	if _, err := queryDatabase("SELECT * FROM non_existent_table"); err != nil {
		// Wrap the error with additional context.
		// fmt.Errorf and the %w verb is used to indicating the error should be
		// wrapped, with the original reference to the error maintained.
		err = fmt.Errorf("failed to query database: %w", err)
		fmt.Println("Error: ", err)
	}
}

func queryDatabase(query string) (*sql.Rows, error) {
	db, err := sql.Open("postgres", "database-dsn")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		// Wrap the error to add context that it occurred during querying.
		return nil, fmt.Errorf("query execution failed: %w", err)
	}

	return rows, nil
}

func main() {
	queryAndProcess()
}
