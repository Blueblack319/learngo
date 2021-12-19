package myDict

import "errors"

// errors
var errNotFound = errors.New("not Found")
var errWordExits = errors.New("that word already exists")

// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errNotFound
	}
	return definition, nil
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExits
	}
	return nil
}

// Update the word in dictionary
