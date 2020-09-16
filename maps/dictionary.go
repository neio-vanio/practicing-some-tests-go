package maps

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotfound not found
var ErrNotfound = errors.New("We could not find the word you are looking for")

// Search in dictionary
func (d Dictionary) Search(serach string) (string, error) {

	def, exist := d[serach]
	if !exist {
		return "", ErrNotfound
	}
	return def, nil
}

// Add word in dictionary
func (d Dictionary) Add(key, word string) {
	d[key] = word
}
