package hello

import (
	"fmt"
	"io"
)

const prefixHelloPortuguese string = "Hello, "
const prefixHelloSpanish string = "Hola, "
const prefixHelloFrench string = "Bonjour, "

// Hello text
func Hello(name string, idioma string) string {
	if name == "" {
		name = "world"
	}

	return prefixGet(idioma) + name
}

func prefixGet(idioma string) (prefix string) {

	switch idioma {
	case "spanish":
		prefix = prefixHelloSpanish
	case "french":
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloPortuguese
	}
	return

}

// SimpleHello for write
func SimpleHello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello, %s", name)
}

func main() {
	fmt.Println(Hello("Vanio", ""))
}
