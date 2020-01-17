package main

import (
	"regexp"

	"github.com/ozgio/strutil"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates random string with given length.
func RandomString(length int) (string, error) {
	return strutil.Random(chars, length)
}

// DefaultProtocol represents http protocol part of url.
const DefaultProtocol = "http://"

// HasProtocol checks if given url has any
// protocol at the beginning.
func HasProtocol(url string) bool {
	validProtocol := regexp.MustCompile(`^[a-zA-Z]*[:][/]{2}[^/]`)
	return validProtocol.Match([]byte(url))
}
