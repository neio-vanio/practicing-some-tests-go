package count

import (
	"fmt"
	"io"
	"time"
)

const writing = "writing"
const pausa = "pausa"

// Sleeper await
type Sleeper interface {
	Sleep()
}

// SleeperSpy with qtd
type SleeperSpy struct {
	Qtd int
}

// SleeperCountOperationsSpy for test
type SleeperCountOperationsSpy struct {
	Calls []string
}

// SleeperDefault for prod
type SleeperDefault struct{}

// Sleep implement interface
func (s *SleeperSpy) Sleep() {
	s.Qtd++
}

// Sleep default for prod
func (s *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep default for test
func (s *SleeperCountOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, pausa)
}

func (s *SleeperCountOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writing)
	return
}

// Decrement a number
func Decrement(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, "Go!")
}
