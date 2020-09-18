package maps

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotfound not found
var (
	ErrNotfound        = errors.New("We could not find the word you are looking for")
	ErrWordExisting    = errors.New("You cannot add the word because it already exists")
	ErrWordNonexistent = errors.New("The word could not be updated because it does not exist")
)

// ErrDictionary of dictionary
type ErrDictionary string

// Error to string
func (e ErrDictionary) Error() string {
	return string(e)
}

// Search in dictionary
func (d Dictionary) Search(serach string) (string, error) {

	def, exist := d[serach]
	if !exist {
		return "", ErrNotfound
	}
	return def, nil
}

// Add word in dictionary
func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotfound:
		d[key] = word
	case nil:
		return ErrWordExisting
	default:
		return err
	}

	return nil
}

// Update word in dictionary
func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotfound:
		return ErrWordNonexistent
	case nil:
		d[key] = word
	default:
		return err
	}

	return nil
}

// Delete key dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
