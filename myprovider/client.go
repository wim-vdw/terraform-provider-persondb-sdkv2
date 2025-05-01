package myprovider

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	CustomDatabase string
}

func (c *Client) initDB() error {
	db, err := sql.Open("sqlite3", c.CustomDatabase)
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS persons (
		person_id TEXT NOT NULL PRIMARY KEY,
		last_name TEXT NOT NULL,
		first_name TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createPerson(personID, lastName, firstName string) error {
	db, err := sql.Open("sqlite3", c.CustomDatabase)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO persons (person_id, last_name, first_name) VALUES (?, ?, ?)", personID, lastName, firstName)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) deletePerson(personID string) error {
	db, err := sql.Open("sqlite3", c.CustomDatabase)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM persons WHERE person_id = ?", personID)
	if err != nil {
		return err
	}
	return nil
}
