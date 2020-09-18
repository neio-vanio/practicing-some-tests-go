package main

import (
	"os"
	"time"

	"github.com/neio-vanio/practicing-some-tests-go/count"
)

func main() {
	sleeper := &count.SleeperConfig{1 * time.Second, time.Sleep}
	count.Decrement(os.Stdout, sleeper)
}
