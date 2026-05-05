package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("Roma é amor."))                   // true
	fmt.Println(isPalindrome("No 'x' in Nixon"))                // true
	fmt.Println(isPalindrome("Hello, World!"))                  // false
	fmt.Println(isPalindrome("A torre da derrota."))            // true
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !isAlphaNumeric(s[left]) {
			left++
		} else if !isAlphaNumeric(s[right]) {
			right--
		} else if toLower(s[left]) != toLower(s[right]) {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
