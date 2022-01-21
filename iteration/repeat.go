package iteration

import "strings"

//Repeat the given letter multiple times
func Repeat(letter string, times int) string {
	var builder strings.Builder

	for i := 0; i < times; i++ {
		builder.WriteString(letter)
	}
	return builder.String()
	// var repeated string
	// for i := 0; i < times; i++ {
	// 	repeated += letter
	// }
	// return repeated
}
