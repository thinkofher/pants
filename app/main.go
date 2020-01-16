package main

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	"github.com/ozgio/strutil"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const KeyLength = 7

var keyNoValueErr = errors.New("db: there is no value assigned to given key")
var keyTakenErr = errors.New("db: key is already taken")

func RandomString(length int) (string, error) {
	return strutil.Random(chars, length)
}

type Entry struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func NewEntry(URL string) (*Entry, error) {
	key, err := RandomString(KeyLength)
	if err != nil {
		return nil, err
	}

	return &Entry{Key: key, URL: URL}, nil
}

type DB map[string]string

func (db DB) SaveEntry(e Entry) error {
	_, ok := db[e.Key]
	if ok {
		return keyTakenErr
	}

	db[e.Key] = e.URL
	return nil
}

func (db DB) GetURL(key string) (string, error) {
	_, ok := db[key]
	if !ok {
		return "", keyNoValueErr
	}

	return db[key], nil
}

func (db DB) GetEntry(key string) (*Entry, error) {
	value, err := db.GetURL(key)
	if err != nil {
		return nil, err
	}

	return &Entry{Key: key, URL: value}, nil
}

type URL struct {
	Value string `json:"value"`
}

func ShortURL(db *DB) echo.HandlerFunc {

	var bindErr = echo.NewHTTPError(http.StatusBadRequest,
		"Failed to parse given data.")
	var shortErr = echo.NewHTTPError(http.StatusInternalServerError,
		"Could not short given url.")
	var luckyErr = echo.NewHTTPError(http.StatusInternalServerError,
		"Could not short given url. Please try again.")

	return func(c echo.Context) error {

		url := &URL{Value: ""}
		if err := c.Bind(url); err != nil {
			return bindErr
		}

		entry, err := NewEntry(url.Value)
		if err != nil {
			return shortErr
		}

		err = db.SaveEntry(*entry)
		if err == keyTakenErr {
			return luckyErr
		} else if err != nil {
			return shortErr
		}

		return c.JSON(http.StatusCreated, entry)
	}
}

func GetShort(db *DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")
		entry, err := db.GetEntry(key)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				"There is no entry behind given key.")
		}

		return c.JSON(http.StatusOK, entry)
	}
}

const DefaultProtocol = "http://"

func RedirectToShort(db *DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Param("url")

		goalURL, err := db.GetURL(url)
		if err == keyNoValueErr {
			return echo.NewHTTPError(http.StatusInternalServerError,
				"There is no value behind given key.")
		}

		if !HasProtocol(goalURL) {
			return c.Redirect(http.StatusTemporaryRedirect, DefaultProtocol+goalURL)
		}

		return c.Redirect(http.StatusTemporaryRedirect, goalURL)
	}
}

// HasProtocol checks if given url has any
// protocol at the beginning.
func HasProtocol(url string) bool {
	validProtocol := regexp.MustCompile(`^[a-zA-Z]*[:][/]{2}[^/]`)
	return validProtocol.Match([]byte(url))
}

func main() {
	// Echo instance
	e := echo.New()

	db := make(DB)

	// Routes
	e.GET("/api/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"value": "Hello World!",
		})
	})
	e.POST("/api/short", ShortURL(&db))
	e.GET("/api/short/:key", GetShort(&db))
	e.GET("/:url", RedirectToShort(&db))

	e.Logger.Fatal(e.Start(":8080"))
}
