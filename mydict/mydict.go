package mydict

import "errors"

// Dictionary
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Word Exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition

	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word) // delete return anything
}
