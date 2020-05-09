package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord("Hello  World"))
	fmt.Println(lengthOfLastWord("Hello"))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("   "))
}
