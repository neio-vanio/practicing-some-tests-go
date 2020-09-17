package hello

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("in french", func(t *testing.T) {
		name := "Vanio"
		result := Hello(name, "french")
		expected := prefixHelloFrench + name
		checkCorrectMessage(t, result, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		name := "Vanio"
		result := Hello(name, "spanish")
		expected := prefixHelloSpanish + name
		checkCorrectMessage(t, result, expected)
	})

	t.Run("expected hello to people", func(t *testing.T) {
		name := "Vanio"
		result := Hello(name, "")
		expected := prefixHelloPortuguese + name
		checkCorrectMessage(t, result, expected)
	})

	t.Run("expected hello world when name not for informed", func(t *testing.T) {
		name := ""
		result := Hello(name, "")
		expected := prefixHelloPortuguese + "world"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("Test simple hello", func(t *testing.T) {
		buffer := bytes.Buffer{}
		SimpleHello(&buffer, "Vanio")

		result := buffer.String()
		expected := "hello, Vanio"

		if result != expected {
			t.Errorf("result '%s' expected '%s'", result, expected)
		}
	})

}
