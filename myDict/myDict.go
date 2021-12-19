package myDict

import "errors"

// errors
var (
	errNotUpdated = errors.New("that word not exists")
	errNotFound   = errors.New("not Found")
	errWordExits  = errors.New("that word already exists")
)

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
func (d Dictionary) Add(word, def string) error {
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
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNotUpdated
	case nil:
		d[word] = definition
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
