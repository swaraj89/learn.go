package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// banner("Go", 15)
	// banner("G😎", 15)

	// s := "G😎"
	// fmt.Println("len:", len(s))

	// for i, r := range s {
	// 	fmt.Println(i, r)
	// }

	// b := s[0]
	// fmt.Printf("%c of type %T\n", b, b)

	// x, y := 1, "1"
	// fmt.Printf("x=%v, y=%v\n", x, y)
	// fmt.Printf("x=%#v, y=%#v\n", x, y)

	// fmt.Printf("%20s", s)

	fmt.Printf("%s is palindrome: %v", "go", isPalindrome("go"))
	fmt.Println()
	fmt.Printf("%s is palindrome: %v", "gog", isPalindrome("gog"))
	fmt.Println()
	fmt.Printf("%s is palindrome: %v", "goog", isPalindrome("goog"))
	fmt.Println()
	fmt.Printf("%s is palindrome: %v", "googl", isPalindrome("googl"))
	fmt.Println()
	fmt.Printf("%s is palindrome: %v", "google", isPalindrome("google"))
	fmt.Println()
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString((text))) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(text string) bool {
	rtext := []rune(text)
	for i := 0; i < len(text)/2; i++ {
		if rtext[i] != rtext[len(rtext)-i-1] {
			return false
		}
	}
	return true
}
