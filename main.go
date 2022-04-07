package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", errors.New("invalid UTF-8 string")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumps over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %s\n", input)
	fmt.Printf("reversed: %q,err: %v\n", rev, revErr)
	fmt.Printf("double reversed: %q,err: %v\n", doubleRev, doubleRevErr)
}
