package main

import "fmt"

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	input := "The quick brown fox jumps over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %s\n", input)
	fmt.Printf("reversed: %s\n", rev)
	fmt.Printf("reversed again: %s\n", doubleRev)
}
