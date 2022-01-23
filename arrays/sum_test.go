package arrays

import (
	"testing"
)

func Test(t *testing.T) {

	t.Run("Sum collection of 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, 8}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %q got %q given %v", want, got, numbers)
		}
	})
}
