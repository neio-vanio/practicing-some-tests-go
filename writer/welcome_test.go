package writer

import (
	"bytes"
	"testing"
)

func TestBeWelcome(t *testing.T) {
	buffer := bytes.Buffer{}
	mesasge := "welcome Vanio"

	BeWelcome(&buffer, mesasge)

	result := buffer.String()
	expected := mesasge

	if result != expected {
		t.Errorf("result: '%s', expected: '%v'", result, expected)
	}
}
