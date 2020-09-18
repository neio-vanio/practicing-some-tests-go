package main

import (
	"os"

	"github.com/neio-vanio/practicing-some-tests-go/count"
)

func main() {
	sleeper := &count.SleeperDefault{}
	count.Decrement(os.Stdout, sleeper)
}
