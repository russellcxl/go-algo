package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

// reverses string using bytes -- which is incorrect
// as it may sometimes lead to malformed characters
// e.g. if you reverse 泃, you get a string that has 3 runes instead of 1
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

func Reverse(s string) (string, error) {

	// without this validation, the string will be converted into an unrecognised rune
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
