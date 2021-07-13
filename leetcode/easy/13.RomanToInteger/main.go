package main

import "fmt"

func RomanToInterger(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	var prev byte
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		add := roman[char]

		if roman[prev] > roman[char] {
			add = -add
		}

		result += add
		prev = char
	}



	return result
}

func main() {
	fmt.Println(RomanToInterger("III"))
}
