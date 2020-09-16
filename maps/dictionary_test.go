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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")

	expected := "this is a test"
	result, err := dictionary.Search("test")
	if err != nil {
		t.Fatal(ErrNotfound, err)
	}

	verifySearch(t, result, expected)
}
