package main

type Entry struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func NewEntry(url string) (*Entry, error) {
	key, err := RandomString(KeyLength)
	if err != nil {
		return nil, err
	}

	return &Entry{Key: key, URL: url}, nil
}
