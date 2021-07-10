package main

// based on : https://leetcode.com/problems/palindrome-number/

// isPalindrome checking our number is palindrome or not
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var rev int

	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}

	return x == rev || x == rev/10
}
