package count

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecrement(t *testing.T) {
	t.Run("decrement with sleep ", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}

		Decrement(buffer, sleeperSpy)

		result := buffer.String()
		expected := "3\n2\n1\nGo!"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}

		if sleeperSpy.Qtd != 4 {
			t.Errorf("no hubo suficiente llamada")
		}
	})

	t.Run("decrement with pause before each print", func(t *testing.T) {
		sleeperCountOperationsSpy := &SleeperCountOperationsSpy{}
		Decrement(sleeperCountOperationsSpy, sleeperCountOperationsSpy)

		expected := []string{
			pausa,
			writing,
			pausa,
			writing,
			pausa,
			writing,
			pausa,
			writing,
		}

		if !reflect.DeepEqual(expected, sleeperCountOperationsSpy.Calls) {
			t.Errorf("expected %v calls, result %v", expected, sleeperCountOperationsSpy.Calls)
		}
	})
}
