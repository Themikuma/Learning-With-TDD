package iteration

import (
	"fmt"
	"testing"
)

const benchmarkString = "this should brick concat"
const benchmarkRepetitions = 100000

func TestRepeat(t *testing.T) {

	formattedAssert := func(t *testing.T, expected string, actual string) {
		t.Helper()
		if actual != expected {
			t.Errorf("exected %q, bot got %q", expected, actual)
		}

	}

	t.Run("repeat a 5  times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"
		formattedAssert(t, expected, actual)
	})

	t.Run("repeat ac 3 times", func(t *testing.T) {
		actual := Repeat("ac", 3)
		expected := "acacac"
		formattedAssert(t, expected, actual)
	})

	t.Run("repeat abc 0 times", func(t *testing.T) {
		actual := Repeat("abc", 0)
		expected := ""
		formattedAssert(t, expected, actual)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(benchmarkString, benchmarkRepetitions)
	}
}

func BenchmarkConcatRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatRepeat(benchmarkString, benchmarkRepetitions)
	}
}

func concatRepeat(letter string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += letter
	}
	return repeated
}

func ExampleRepeat() {
	repeated := Repeat("abc", 3)
	fmt.Println(repeated)
	//Output: abcabcabc
}
