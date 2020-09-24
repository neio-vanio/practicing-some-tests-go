package reflection

import "testing"

func TestWalks(t *testing.T) {
	expected := "Vanio"
	var result []string

	x := struct {
		Name string
	}{expected}

	walks(x, func(arg string) {
		result = append(result, arg)
	})

	if result[0] != expected {
		t.Errorf("result '%v', expected '%v'", result[0], expected)
	}
}
