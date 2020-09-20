package request

import (
	"fmt"
	"net/http"
	"time"
)

func requestWithSelect(a, b string, timeLimit time.Duration) (string, error) {
	select {
	case <-pingRequest(a):
		return a, nil
	case <-pingRequest(b):
		return b, nil
	case <-time.After(timeLimit):
		return "", fmt.Errorf("time slow for %s, %s", a, b)
	}
}

func timeResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func pingRequest(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
