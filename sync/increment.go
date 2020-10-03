package sync

import (
	"sync"
)

// Increment is an accountant
type Increment struct {
	mu    sync.Mutex
	value int
}

// Add + 1
func (i *Increment) Add() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.value++
}

// Value get value
func (i *Increment) Value() int {
	return i.value
}
