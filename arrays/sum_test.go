package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum collection of 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, 8}

		got := Sum(numbers)
		want := 19

		if got != want {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Helper()
	var formattedAssert = func(t testing.TB, got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	}
	
	t.Run("multiple collections with elements", func(t *testing.T) {
		want := []int{20, 10}
		got := SumAll([]int{4, 6, 10}, []int{2, 3, 5})
		formattedAssert(t, got, want)
	})

	t.Run("multiple collections with no elements", func(t *testing.T) {
		want := []int{0, 0}
		got := SumAll([]int{}, []int{})
		formattedAssert(t, got, want)
	})

	t.Run("single collection with no elements", func(t *testing.T) {
		want := []int{0}
		got := SumAll([]int{})
		formattedAssert(t, got, want)
	})
}
