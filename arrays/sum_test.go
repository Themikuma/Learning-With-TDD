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

func arrayAssertHelper(t *testing.T, want []int, got []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSumAll(t *testing.T) {

	t.Run("multiple collections with elements", func(t *testing.T) {
		want := []int{20, 10}
		got := SumAll([]int{4, 6, 10}, []int{2, 3, 5})
		arrayAssertHelper(t, got, want)
	})

	t.Run("multiple collections with no elements", func(t *testing.T) {
		want := []int{0, 0}
		got := SumAll([]int{}, []int{})
		arrayAssertHelper(t, got, want)
	})

	t.Run("single collection with no elements", func(t *testing.T) {
		want := []int{0}
		got := SumAll([]int{})
		arrayAssertHelper(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	tests := []struct {
		name         string
		numbersToSum [][]int
		want         []int
	}{
		{
			name: "multiple collections with elements",
			numbersToSum: [][]int{
				{3, 1, 2},
				{15, 2, 3}},
			want: []int{3, 5}},
		{
			name: "multiple collections with no elements",
			numbersToSum: [][]int{
				{},
				{}},
			want: []int{0, 0}},
		{
			name: "single collection with single element",
			numbersToSum: [][]int{
				{42}},
			want: []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAllTails(tt.numbersToSum...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAllTails() = %v, want %v", got, tt.want)
			}
		})
	}
}
