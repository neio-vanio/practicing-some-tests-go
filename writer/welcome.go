package writer

import (
	"fmt"
	"io"
)

// BeWelcome is used to write a message to the user
func BeWelcome(writer io.Writer, message string) {
	fmt.Fprintf(writer, message)
}
