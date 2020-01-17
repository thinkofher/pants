package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

// KeyLength represents default length of keys
// used to store values of target urls and
// redirect clients.
const KeyLength = 7

var keyNoValueErr = errors.New("db: there is no value assigned to given key")
var keyTakenErr = errors.New("db: key is already taken")

// DB interface specifies methods for inserting
// urls to database and retrieving them by keys.
type DB interface {
	// Saves Entry in database.
	SaveEntry(e Entry) error
	// Returns url assigned to given key.
	GetURL(key string) (string, error)
	// Returns full Entry (specified in API) assigned
	// to given key.
	GetEntry(key string) (*Entry, error)
}

type redisDB struct {
	client *redis.Client
}

func (db redisDB) SaveEntry(e Entry) error {
	err := db.client.Set(e.Key, e.URL, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db redisDB) GetURL(key string) (string, error) {
	val, err := db.client.Get(key).Result()
	if err != nil {
		return "", keyNoValueErr
	}

	return val, nil
}

func (db redisDB) GetEntry(key string) (*Entry, error) {
	val, err := db.GetURL(key)
	if err != nil {
		return nil, err
	}

	return &Entry{Key: key, URL: val}, nil
}

type URL struct {
	Value string `json:"value"`
}

func ShortURL(db DB) echo.HandlerFunc {
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

func GetShort(db DB) echo.HandlerFunc {
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

func RedirectToShort(db DB) echo.HandlerFunc {
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

func main() {
	// Redis initialization
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Printf("Redis started at '%s'.\n", client.Options().Addr)
	db := redisDB{client: client}

	// Echo instance
	e := echo.New()

	// Routes
	e.POST("/api/short", ShortURL(&db))
	e.GET("/api/short/:key", GetShort(&db))
	e.GET("/:url", RedirectToShort(&db))

	e.Logger.Fatal(e.Start(":8080"))
}
