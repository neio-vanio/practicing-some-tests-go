package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is a test"

		verifySearch(t, result, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		verifyError(t, err, ErrNotfound)
	})

}

func verifySearch(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result %s, expected %s search %s", result, expected, "test")
	}
}

func verifyError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result error %s, expected %s", result, expected)
	}
}

func verifyDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found added word:", err)
	}

	if definition != result {
		t.Errorf("result '%s',  expected '%s'", result, definition)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is a test")

		expected := "this is a test"
		result, err := dictionary.Search("test")

		verifyError(t, err, nil)
		verifySearch(t, result, expected)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add("test", "this is a test")

		verifyError(t, err, ErrWordExisting)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing definition", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		verifyDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new definition", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		verifyError(t, err, ErrWordNonexistent)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "this is a test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	verifyError(t, err, ErrNotfound)
}
