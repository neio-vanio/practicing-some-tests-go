package sync

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	t.Run("Increment added 3 times - non-competitor", func(t *testing.T) {
		increment := &Increment{}
		increment.Add()
		increment.Add()
		increment.Add()

		verifyValue(t, increment, 3)
	})

	t.Run("Increment added 3 times - competitor", func(t *testing.T) {
		increment := &Increment{}
		expected := 1000

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func(w *sync.WaitGroup) {
				increment.Add()
				wg.Done()
			}(&wg)
		}
		wg.Wait()

		verifyValue(t, increment, expected)
	})
}

func verifyValue(t *testing.T, increment *Increment, expected int) {
	t.Helper()
	if increment.Value() != expected {
		t.Errorf("result %d, expected %d", increment.Value(), expected)
	}
}
