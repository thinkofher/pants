package main

import "database/sql"

type dbWrapper struct {
	client *sql.DB
}

func InitDB(file string) (*dbWrapper, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	statement := `
	CREATE TABLE IF NOT EXISTS urls
		(id INTEGER PRIMARY KEY, key TEXT, url TEXT);
	`
	_, err = db.Exec(statement)
	if err != nil {
		return nil, err
	}

	return &dbWrapper{client: db}, nil
}

func (db dbWrapper) SaveEntry(e Entry) error {
	statement := `
	INSERT INTO urls (key, url)
	VALUES
		($1, $2);
	`
	_, err := db.client.Exec(statement, e.Key, e.URL)
	if err != nil {
		return err
	}

	return nil
}

func (db dbWrapper) GetURL(key string) (string, error) {
	var ans string

	statement := `
	SELECT
		url
	FROM
		urls
	WHERE
		key = $1;
	`
	err := db.client.QueryRow(statement, key).Scan(&ans)
	if err != nil {
		return "", keyNoValueErr
	}

	return ans, nil
}

func (db dbWrapper) GetEntry(key string) (*Entry, error) {
	val, err := db.GetURL(key)
	if err != nil {
		return nil, err
	}

	return &Entry{Key: key, URL: val}, nil
}
