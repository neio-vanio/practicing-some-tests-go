package main

import (
	"net/http"
	"os"

	writer "github.com/neio-vanio/practicing-some-tests-go/writer"
)

func main() {
	writer.BeWelcome(os.Stdout, "welcome Vanio")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writer.BeWelcome(w, "welcome Vanio")
	})
	http.ListenAndServe(":8777", nil)
}
