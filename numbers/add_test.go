package numbers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 4
	result := Add(2, 2)

	if result != expected {
		t.Errorf("expected '%d', result '%d'", expected, result)
	}
}

func ExampleAdd() {
	sum := Add(7, 5)
	fmt.Println(sum)
}
