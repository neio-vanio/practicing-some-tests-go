package competition

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func mockCheckerURI(URI string) bool {
	if URI == "htt://neio.network" {
		return false
	}
	return true
}

func TestVerifyURI(t *testing.T) {
	URIS := []string{
		"https://neio.network",
		"https:google.com",
		"htt://neio.network",
	}

	expected := map[string]bool{
		"https://neio.network": true,
		"https:google.com":     true,
		"htt://neio.network":   false,
	}

	result := VerifyURI(mockCheckerURI, URIS)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, result %v", expected, result)
	}
}

func slowCheckerURI(URI string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerifyURI(t *testing.B) {
	URIS := make([]string, 100)
	for i := 0; i < len(URIS); i++ {
		URIS[i] = fmt.Sprintf("http://%s.com.br", strconv.Itoa(i))
	}

	for i := 0; i < t.N; i++ {
		VerifyURI(slowCheckerURI, URIS)
	}
}
