package iteration

import (
	"reflect"
	"testing"
)

var verifySum = func(t *testing.T, result, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result '%v' expected '%v'", result, expected)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("result %d, want %d, dado %v", result, expected, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	verifySum(t, result, expected)
}

func TestSumOfEverythingElse(t *testing.T) {
	t.Run("Does the sum of some slices", func(t *testing.T) {
		result := SumOfEverythingElse([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		verifySum(t, result, expected)
	})

	t.Run("Makes the sums of some empty slices safely", func(t *testing.T) {
		result := SumOfEverythingElse([]int{}, []int{0, 9})
		expected := []int{0, 9}
		verifySum(t, result, expected)
	})

}
