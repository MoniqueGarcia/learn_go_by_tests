package iteration

import "strings"

func Repeat(character string, repeat int) string {
	// Using for loop
	// var result string
	// for i:=0; i < repeat; i++ {
	// 	result += character
	// }
	// return result

	// Using Repeat() func from strings package
	return strings.Repeat(character, repeat)
}